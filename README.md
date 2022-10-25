"Content-Type", "application/vnd.kafka.binary.v2+json")
"X-Real-Ip", "10.48.5.59")
"X-Original-Uri", "/topics/000-0.sap-erp.db.operations.orders05.0")
"X-Original-Method", "POST")
"X-Service", "kafka-rest")
"Authorization", "Basic c2FwOnNlQzIzc0JGanV0azg5TnY=")


auth:
  prefix: /topics/
  urlvalidreg: ^\d{3}-\d(-\d{3}-\d)?\.[a-z0-9-]+\.(db|cdc|cmd|sys|log|tmp)\.[a-z0-9-.]+\.\d+$
  acl:
  - path: ^888-8\.example\.db\.
    users:
    - avro-user
    methods:
    - POST
    contenttype:
    - application/vnd.kafka.avro.v2+json
