package continuous_evaluation_go

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
)

func SendKPI(kpi, value string) {
	if ctl := os.Getenv("CONTEVAL_CONTROLLER"); len(ctl) > 0 {
		http.Get(fmt.Sprintf("%s?commit=%s&kpi=%s&value=%s",
			ctl, os.Getenv("CONTEVAL_COMMIT"), kpi, value))
	} else {
		log.Printf("Not started by Calla; no sending %s\n", kpi)
	}
}

func TestExampleUnitEvaluation(t *testing.T) {
	SendKPI("AnExampleKPI", "10")
}
