# Monitor Services

Run queue server

```bash
bash .bin/serve_queue -mongoUrl $MONGO_URL
``` 

Run queue event dispatcher
```bash
bash .bin/dispatch_queue -mongoUrl $MONGO_URL
```

Run US bills server 
```bash
bash .bin/serve_us_bills -mongoUrl $MONGO_URL -apiKey $GOVINFO_API_KEY
```

Run US bills fetch
```bash
bash .bin/fetch_us_bills -mongoUrl $MONGO_URL -apiKey $GOVINFO_API_KEY 
```