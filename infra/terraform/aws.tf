provider "aws" {
  region                      = "us-east-2"
  access_key                  = "fake"
  secret_key                  = "fake"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
  endpoints {
    dynamodb = "http://localhost:4566"
    sqs      = "http://localhost:4566"
    sns      = "http://localhost:4566"
  }
}

resource "aws_dynamodb_table" "link_redirect" {
  name           = "link_redirect"
  billing_mode   = "PROVISIONED"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "PK"
  range_key      = "SK"

  attribute {
    name = "PK"
    type = "S"
  }

  attribute {
    name = "SK"
    type = "S"
  }
}

resource "aws_sns_topic" "url_management_short_url" {
  name       = "url_management_short_url_topic.fifo"
  fifo_topic = true
}

resource "aws_sqs_queue" "url_redirect_short_url" {
  name       = "url_redirect_short_url_queue.fifo"
  fifo_queue = true
}

resource "aws_sns_topic_subscription" "url_management_short_url_url_redirect" {
  topic_arn = aws_sns_topic.url_management_short_url.arn
  protocol  = "sqs"
  endpoint  = aws_sqs_queue.url_redirect_short_url.arn
}
