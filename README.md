# Go Smarty

Go Smarty is a simple web application built with Go and the Gin framework. It provides functionality to upload files to a cloud storage bucket and record information about the uploaded files in a PostgreSQL database.

## Setup

Before running the application, make sure you have Go installed on your system. You also need to set up a Google Cloud Storage bucket and a PostgreSQL database.

## Usage
You can upload files by sending a POST request to the /upload endpoint with a file parameter containing the file to be uploaded.