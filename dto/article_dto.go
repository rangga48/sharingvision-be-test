package dto

import (
	"errors"
	"fmt"
	"strings"
)

type ArticleRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

type ArticleResponse struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	CreatedDate string `json:"created_date,omitempty"`
	UpdatedDate string `json:"updated_date,omitempty"`
	Status      string `json:"status"`
}

func (req *ArticleRequest) Validate() error {
	var errMsgs []string

	trimmedTitle := strings.TrimSpace(req.Title)
	if trimmedTitle == "" {
		errMsgs = append(errMsgs, "Title is required")
	} else if len(trimmedTitle) < 20 {
		errMsgs = append(errMsgs, "Title minimal 20 karakter")
	}

	trimmedContent := strings.TrimSpace(req.Content)
	if trimmedContent == "" {
		errMsgs = append(errMsgs, "Content is required")
	} else if len(trimmedContent) < 200 {
		errMsgs = append(errMsgs, "Content minimal 200 karakter")
	}

	trimmedCategory := strings.TrimSpace(req.Category)
	if trimmedCategory == "" {
		errMsgs = append(errMsgs, "Category is required")
	} else if len(trimmedCategory) < 3 {
		errMsgs = append(errMsgs, "Category minimal 3 karakter")
	}

	validStatuses := map[string]bool{
		"publish": true,
		"draft":   true,
		"thrash":  true,
		"Publish": true,
		"Draft":   true,
		"Thrash":  true,
	}

	trimmedStatus := strings.TrimSpace(req.Status)
	if trimmedStatus == "" {
		errMsgs = append(errMsgs, "Status is required")
	} else if !validStatuses[trimmedStatus] {
		errMsgs = append(errMsgs, "Status harus memilih antara 'publish', 'draft', atau 'thrash'")
	}

	if len(errMsgs) > 0 {
		return errors.New(fmt.Sprintf("Validation Error: %s", strings.Join(errMsgs, "; ")))
	}

	return nil
}
