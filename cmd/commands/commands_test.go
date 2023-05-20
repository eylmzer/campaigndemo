package commands

import (
	"testing"

	"github.com/eylmzer/campaingdemo/pkg/campaingscenario"
)

func TestExecuteCommand(t *testing.T) {
	cs := campaingscenario.NewCampaingScenario(nil)

	// Test case 1: Create Product
	cmd := "create_product P1 10.0 100"
	expectedOutput := "Product created; code P1, price 10.00, stock 100"
	output, err := ExecuteCommand(cmd, cs)
	if err != nil {
		t.Errorf("Error executing command: %v", err)
	}
	if output != expectedOutput {
		t.Errorf("Expected output '%s', but got '%s'", expectedOutput, output)
	}

	// Test case 2: Create Order
	cmd = "create_order P1 5"
	expectedOutput = "Order created; product P1, quantity 5, price 10.00"
	output, err = ExecuteCommand(cmd, cs)
	if err != nil {
		t.Errorf("Error executing command: %v", err)
	}
	if output != expectedOutput {
		t.Errorf("Expected output '%s', but got '%s'", expectedOutput, output)
	}

	// Test case 3: Create Campaign
	cmd = "create_campaign C1 P1 24 20.0 100"
	expectedOutput = "Campaign created; name C1, product P1, duration 24, limit 20.00, target sales count 100"
	output, err = ExecuteCommand(cmd, cs)
	if err != nil {
		t.Errorf("Error executing command: %v", err)
	}
	if output != expectedOutput {
		t.Errorf("Expected output '%s', but got '%s'", expectedOutput, output)
	}

	// Test case 4: Get Product Info
	cmd = "get_product_info P1"
	expectedOutput = "Product P1 info; price 10.00, stock 100"
	output, err = ExecuteCommand(cmd, cs)
	if err != nil {
		t.Errorf("Error executing command: %v", err)
	}
	if output != expectedOutput {
		t.Errorf("Expected output '%s', but got '%s'", expectedOutput, output)
	}

	// Test case 5: Increase Time
	cmd = "increase_time 6"
	expectedOutput = "Time is 06:00"
	output, err = ExecuteCommand(cmd, cs)
	if err != nil {
		t.Errorf("Error executing command: %v", err)
	}
	if output != expectedOutput {
		t.Errorf("Expected output '%s', but got '%s'", expectedOutput, output)
	}

	// Test case 6: Get Campaign Info
	cmd = "get_campaign_info C1"
	expectedOutput = "Campaign C1 info; Status Active, Target Sales 100, Total Sales 0, Turnover 0.00, Average Item Price 0.00"
	output, err = ExecuteCommand(cmd, cs)
	if err != nil {
		t.Errorf("Error executing command: %v", err)
	}
	if output != expectedOutput {
		t.Errorf("Expected output '%s', but got '%s'", expectedOutput, output)
	}

	// Test case 7: Invalid Command
	cmd = "invalid_command"
	expectedOutput = "invalid command"
	_, err = ExecuteCommand(cmd, cs)
	if err == nil || err.Error() != expectedOutput {
		t.Errorf("Expected error '%s', but got '%v'", expectedOutput, err)
	}
}
