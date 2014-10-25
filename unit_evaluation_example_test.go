package continuous_evaluation_go

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestExampleUnitEvaluation(t *testing.T) {
	if ctl := os.Getenv("CONTEVAL_CONTROLLER"); len(ctl) > 0 {
		http.Get(fmt.Sprintf("%s?commit=%s&kpi=AnExampleKPI&value=10",
			ctl, os.Getenv("CONTEVAL_COMMIT")))
	} else {
		fmt.Println("Not started by Calla; skip TestExampleUnitEvaluation")
	}
}
