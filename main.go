package main

import (
   "context"
   "fmt"
   "log"
  
   aiplatform "cloud.google.com/go/aiplatform/apiv1"
   aiplatformpb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
   "google.golang.org/api/option"
   "google.golang.org/protobuf/types/known/structpb"
)

func main() {
   ctx := context.Background()
   c, err := aiplatform.NewPredictionClient(ctx, option.WithCredentialsFile("/Users/pwm/Documents/cloud-credentials/gcp/spyderviews-sa.json"), option.WithEndpoint("us-central1-aiplatform.googleapis.com:443"))
   if err != nil {
      log.Fatalf("Error 1: %v", err)
   }
   defer c.Close()
 
   params, err := structpb.NewValue(map[string]interface{}{
      "temperature":     0.2,
      "maxOutputTokens": 1024,
      "topK":            40,
   })
   if err != nil {
      log.Fatalf("Error 2: %v", err)
   }

   instance, err := structpb.NewValue(map[string]interface{}{
      "prompt": fmt.Sprintf("what's the date today?"),
   })
   if err != nil {
      log.Fatalf("Error 3: %v", err)
   }
  
   req := &aiplatformpb.PredictRequest{
      Endpoint:   "projects/speedy-victory-336109/locations/us-central1/publishers/google/models/text-unicorn@001",
      Instances:  []*structpb.Value{instance},
      Parameters: params,
   }
   resp, err := c.Predict(ctx, req)
   if err != nil {
      log.Fatalf("Error 4: %v", err)
   }
  
   respMap := resp.Predictions[0].GetStructValue().AsMap()
   fmt.Printf("resp: %v", respMap["content"])
}