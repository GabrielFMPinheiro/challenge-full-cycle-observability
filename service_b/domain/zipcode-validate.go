package domain

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"os"

	"go.opentelemetry.io/otel"
)

func ZipcodeValidate(ctx context.Context, w http.ResponseWriter, marshalled []byte) error {
	ctx, span := otel.GetTracerProvider().Tracer("service_a").Start(ctx, "service_a")
	defer span.End()

	req, err := http.NewRequestWithContext(ctx, "POST",
		os.Getenv("SERVICE_A_URL")+"/zipcode", bytes.NewBuffer(marshalled))

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	client := http.Client{}

	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid zipcode")
	}

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
