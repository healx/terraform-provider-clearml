package client

import (
	"context"
)

func (c *ClearMLClient) CreateQueue(ctx context.Context, name string, tags []string) (string, error) {

	payload := map[string]interface{}{
		"name": name,
		"tags": tags,
	}

	res, err := c.request(ctx, "POST", "/queues.create", nil, payload)

	if err != nil {
		return "", err
	}

	dataObj := res.(map[string]interface{})
	queueObj := dataObj["data"].(map[string]interface{})

	return queueObj["id"].(string), nil
}

func (c *ClearMLClient) GetQueue(ctx context.Context, id string) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"queue": id,
	}

	res, err := c.request(ctx, "POST", "/queues.get_by_id", nil, payload)

	if err != nil {
		return nil, err
	}

	dataObj := res.(map[string]interface{})
	queueObj := dataObj["data"].(map[string]interface{})

	return queueObj["queue"].(map[string]interface{}), nil
}

func (c *ClearMLClient) UpdateQueue(ctx context.Context, id string, name string, tags []string) error {
	payload := map[string]interface{}{
		"queue": id,
		"name": name,
		"tags": tags,
	}

	_, err := c.request(ctx, "POST", "/queues.update", nil, payload)

	if err != nil {
		return err
	}

	return nil
}

func (c *ClearMLClient) DeleteQueue(ctx context.Context, id string) error {
	payload := map[string]interface{}{
		"queue": id,
	}

	_, err := c.request(ctx, "POST", "/queues.delete", nil, payload)

	if err != nil {
		return err
	}

	return nil
}