package service

import (
	"errors"
	"fmt"
	"server-slug/database"
	"net/url"
)

func GetAlllinks() ([]DataCreate, error) {
	records, err := database.LoadRecords()
	if err != nil {
		return nil, fmt.Errorf("Ошибка при загрузке ссылок%w", err)
	}
	items := make([]DataCreate, 0, len(records))
	for _, record := range records {
		items = append(items, DataCreate{
			URL:  record.URL,
			Slug: record.Slug,
		})
	}
	return items, nil
}

func GetLinkBySlug(slug string) (*Data, error) {
	records, err := database.LoadRecords()
	if err != nil {
		return nil, fmt.Errorf("Ошибка при загрузке данных %w", err)
	}

	for _, record := range records {
		if record.Slug == slug {
			return &Data{
				ID:   record.ID,
				Slug: record.Slug,
				URL:  record.URL,
			}, nil

		}
	}
	return nil, errors.New("Данные не найдены")
}

func PostLink(URL, Slug string) (*Data, error) {

	records, err := database.LoadRecords()
	if err != nil {
		return nil, fmt.Errorf("Ошибка при загрузке данных %w", err)
	}

	newID := generateID(records)

	data := &Data{
		ID:   newID,
		URL:  URL,
		Slug: Slug,
	}

	newData := database.Record{
		ID:   data.ID,
		URL:  data.URL,
		Slug: data.Slug,
	}
	records = append(records, newData)
	if err := database.SaveRecords(records); err != nil {
		return nil, fmt.Errorf("Ошибка при загрузке данных %w", err)
	}
	return data, nil
}

func PatchLinkById(id int, url, slug string) (*Data, error) {
	records, err := database.LoadRecords()
	if err != nil {
		return nil, fmt.Errorf("Ошибка при загрузке данных %w", err)
	}

	found := false
	for i, record := range records {
		if record.ID == id {
			records[i].Slug = slug
			records[i].URL = url
			found = true
			break
		}
	}

	if !found {
		return nil, errors.New("Slug не найден")
	}
	if err := database.SaveRecords(records); err != nil {
		return nil, fmt.Errorf("Ошибка при загрузке данных %w", err)
	}

	return &Data{
		ID:   id,
		Slug: slug,
		URL:  url,
	}, nil
}

func DeleteLinkBySlug(slug string) error {
	records, err := database.LoadRecords()
	if err != nil{
		return fmt.Errorf("Ошибка при зашрузке данных %w", err)
	}

	found := false
	var newRecord []database.Record
	for _, record := range records{
		if record.Slug == slug{
			newRecord = append(newRecord, record)
		} else {found = true }
	}
	
	if !found{
		return errors.New("Ошибка при удалении")
	}
	if err := database.SaveRecords(records); err != nil{
		return fmt.Errorf("Ошибка при загрузке данных %w", err)
	}

	return nil
}

func DeleteLinkByID(id int) error {
	records, err := database.LoadRecords()
	if err != nil{
		return fmt.Errorf("Ошибка при зашрузке данных %w", err)
	}

	found := false
	var newRecord []database.Record
	for _, record := range records{
		if record.ID == id{
			newRecord = append(newRecord, record)
		} else {found = true }
	}
	
	if !found{
		return errors.New("Ошибка при удалении")
	}
	if err := database.SaveRecords(records); err != nil{
		return fmt.Errorf("Ошибка при загрузке данных %w", err)
	}

	return nil
}

func ValidateAndNormalizeURL(rawURL string) (string, error) {
    // Парсим URL
    parsed, err := url.ParseRequestURI(rawURL)
    if err != nil {
        // Пробуем более мягкий парсинг
        parsed, err = url.Parse(rawURL)
        if err != nil {
            return "", fmt.Errorf("invalid URL: %w", err)
        }
    }
    
    // Если нет схемы, добавляем https
    if parsed.Scheme == "" {
        parsed.Scheme = "https"
    }
    
    // Проверяем схему
    if parsed.Scheme != "http" && parsed.Scheme != "https" {
        return "", fmt.Errorf("only http/https URLs are allowed")
    }
    
    // Проверяем, что есть хост
    if parsed.Host == "" {
        return "", fmt.Errorf("missing host in URL")
    }
    
    // Нормализуем URL
    normalizedURL := parsed.String()
	    
    return normalizedURL, nil
}
// func Redirrect(slug string) error{
// 	records, err := database.LoadRecords()
// 	if err != nil{
// 		return fmt.Errorf("Ошибка при загрузке данных %w", err)
// 	}

// 	for _, record := range records{
// 		if record.Slug == slug{
			
// 		}
// 	}
// }

func generateID(records []database.Record) int{
    if len(records) == 0 {
        return 1
    }

    maxID := 0         
    for _, record := range records {
        if record.ID < 0 {
            continue
        }
        
        if record.ID > maxID {
            maxID = record.ID
        }
    }

    return maxID + 1
}  

func SlugAlert(slug string) bool{
	slugs, _ := database.LoadRecords()
	for _, value := range slugs{
		if value.Slug == slug{
			return true
		} 
	}
	return  false
}


func Redirect(redir string) (error, string) {
	r, _ := database.LoadRecords()
	for _, value := range r{
		if value.Slug == redir{
			return nil, value.URL
		}
	}
	return errors.New("Нет такого slug"), ""
}