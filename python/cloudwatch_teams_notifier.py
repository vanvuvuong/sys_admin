# TODO: run when SNS trigger
# Send log error from AWS Cloudwatch subscription filter
import base64
import json
import gzip
import urllib3

CHANNEL_WEBHOOK_URL = ""
# payload = {
#     "messageType": "DATA_MESSAGE",
#     "owner": "776326220523",
#     "logGroup": "cwatch-lgroup-xx-01",
#     "logStream": "logstream",
#     "subscriptionFilters": ["error-filter"],
#     "logEvents": [
#         {
#             "id": "38129764694231871733433282340191055885221457781086027776",
#             "timestamp": 1709797782755,
#             "message": "Error message",
#         },
#         {
#             "id": "38129764786333949403364755914733572351259190800785145857",
#             "timestamp": 1709797786885,
#             "message": "Error message",
#         },
#     ],
# }

http = urllib3.PoolManager()
region = "ap-northeast-1"


def get_error_message(event):
    message = "" or event.get("awslogs", {}).get("data", None)
    return message


def decode_message(message):
    compressed_payload = base64.b64decode(message)
    uncompressed_payload = gzip.decompress(compressed_payload)
    payload = json.loads(uncompressed_payload)
    return payload


def convert_message(payload):
    return ""


def prepare_message(payload):
    url = f"https://{region}.console.aws.amazon.com/cloudwatch/home?region={region}#logsV2:log-groups/log-group/{payload['logGroup']}/log-events/{payload['logStream']}"
    return url


def send_to_teams(text: str, detail_url: str):
    msg = {
        "summary": "Summary",
        "title": "Job title",
        "text": text,
        "potentialAction": [
            {
                "@type": "OpenUri",
                "name": "Detail",
                "targets": [{"os": "default", "uri": detail_url}],
            }
        ],
    }
    encoded_msg = json.dumps(msg).encode("utf-8")
    response = http.request("POST", CHANNEL_WEBHOOK_URL, body=encoded_msg)
    print(
        f"""
        "status_code": {response.status},
        "response": {response.data}
    """
    )


def lambda_handler(event, context):
    message = get_error_message(event)
    if message:
        print(message)
        payload = decode_message(message)
        message_content = convert_message(payload)
        url = prepare_message(payload)
        send_to_teams(message_content, url)
