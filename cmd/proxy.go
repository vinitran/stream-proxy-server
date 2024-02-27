package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grafov/m3u8"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (v *V1Router) m3u8Handler(c *gin.Context) {
	urlQuery := c.Query("url")
	if urlQuery == "" {
		responseFailureWithMessage(c, "url is invalid")
		return
	}

	parsedURL, err := url.Parse(urlQuery)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	baseURL := fmt.Sprintf("%s://%s%s", parsedURL.Scheme, parsedURL.Host, strings.TrimRight(parsedURL.Path, "/"))

	client := &http.Client{}

	req, err := http.NewRequest("GET", urlQuery, nil)
	if err != nil {
		responseErrInternalServerErrorWithDetail(c, "request is invalid")
		return
	}

	req.Header.Set("Referer", REFERER)

	resp, err := client.Do(req)
	if err != nil {
		responseErrInternalServerErrorWithDetail(c, "request is invalid")
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		responseErrInternalServerErrorWithDetail(c, fmt.Sprintf("HTTP request failed with status code: %d", resp.StatusCode))
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		responseErrInternalServerErrorWithDetail(c, "body is invalid")
		return
	}

	p, listType, err := m3u8.DecodeFrom(bufio.NewReader(bytes.NewReader(body)), true)
	if err != nil {
		panic(err)
	}

	switch listType {
	case m3u8.MEDIA:
		mediapl := p.(*m3u8.MediaPlaylist)
		for i := 0; i < len(mediapl.Segments); i++ {
			if mediapl.Segments[i] != nil {
				mediapl.Segments[i].URI = fmt.Sprintf("%s?url=%s", mediapl.Segments[i].URI, baseURL)
			}
		}
		responseSuccessWithData(c, resp.Header.Get("Content-Type"), []byte(mediapl.String()))

	default:
		responseErrInternalServerError(c)
	}
}

func (v *V1Router) hlsHandler(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		responseFailureWithMessage(c, "url is invalid")
		return
	}

	hlsId := c.Param("hls_id")
	if hlsId == "" {
		responseFailureWithMessage(c, "hls is invalid")
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", url, hlsId), nil)
	if err != nil {
		responseErrInternalServerErrorWithDetail(c, "request is invalid")
		return
	}
	req.Header.Set("Referer", REFERER)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		responseErrInternalServerErrorWithDetail(c, "request is invalid")
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		responseErrInternalServerErrorWithDetail(c, fmt.Sprintf("HTTP request failed with status code: %d", resp.StatusCode))
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		responseErrInternalServerErrorWithDetail(c, "body is invalid")
		return
	}

	responseSuccessWithData(c, resp.Header.Get("Content-Type"), body)
}
