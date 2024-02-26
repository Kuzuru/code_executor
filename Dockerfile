FROM python:3.11-slim

# Отдельный user для запуска кода
RUN useradd -rm -d /home/student -s /bin/bash -u 1001 student

# Для компиляции и запуска кода на GoLang
RUN apt-get update && \
apt-get install -y -q --no-install-recommends golang sudo && \
	rm -rf /var/lib/apt/lists/

RUN pip install --no-cache-dir --upgrade pip

WORKDIR /agent

COPY launch.py /agent/launch.py
COPY tempfile_helper.py /agent/tempfile_helper.py
COPY requirements.txt /agent/requirements.txt

RUN mkdir -p /home/student/tests

RUN pip install --no-cache-dir -r requirements.txt
