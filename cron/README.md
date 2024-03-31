# Cron

```bash
python3 bill_urls.py | tee -a billing-urls.txt | python3 bills.py | tee -a bills.txt
```