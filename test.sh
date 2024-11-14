json=$(curl -X POST -H "Content-Type: application/json" -d '{"retailer": "CVS123;&*(#$%456", "purchaseDate": "2020-01-01", "purchaseTime": "15:59", "items": [{"id": "1", "shortDescription": "an item1", "price": "10000000"},{"id": "2", "shortDescription": "   123456     ", "price": "101"},{"id": "3", "shortDescription": "123456   ", "price": "100"}], "total": "100.00"}' http://localhost:8080/receipts/process)
echo $json
id=$(echo $json | jq -r '.id')
curl "http://localhost:8080/receipts/$id/points"
echo "Expected 146 points"

json=$(curl -X POST -H "Content-Type: application/json" -d '{"retailer":"Target","purchaseDate":"2022-01-01","purchaseTime":"13:01","items":[{"shortDescription":"Mountain Dew 12PK","price":"6.49"},{"shortDescription":"Emils Cheese Pizza","price":"12.25"},{"shortDescription":"Knorr Creamy Chicken","price":"1.26"},{"shortDescription":"Doritos Nacho Cheese","price":"3.35"},{"shortDescription":"   Klarbrunn 12-PK 12 FL OZ  ","price":"12.00"}],"total":"35.35"}' http://localhost:8080/receipts/process)
id=$(echo $json | jq -r '.id')
curl "http://localhost:8080/receipts/$id/points"
echo "Expected 28 points"

json=$(curl -X POST -H "Content-Type: application/json" -d '{"retailer":"M&M Corner Market","purchaseDate":"2022-03-20","purchaseTime":"14:33","items":[{"shortDescription":"Gatorade","price":"2.25"},{"shortDescription":"Gatorade","price":"2.25"},{"shortDescription":"Gatorade","price":"2.25"},{"shortDescription":"Gatorade","price":"2.25"}],"total":"9.00"}' http://localhost:8080/receipts/process)
id=$(echo $json | jq -r '.id')
curl "http://localhost:8080/receipts/$id/points"
echo "Expected 109 points"