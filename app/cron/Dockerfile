FROM python:3.10.0-slim

RUN apt update
RUN apt install -y cron
RUN pip install pymongo python-dotenv requests
RUN (crontab -l; echo "*/30 * * * * /usr/local/bin/python /home/app/cron/get_word_api.py") | crontab -
CMD cron -f
