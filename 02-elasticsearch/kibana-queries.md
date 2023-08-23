# Kibana Example Queries

```
# Cluster info
GET _cluster/health

# Node info
GET _nodes/stats

# Create index
PUT cimri_product_index

# Check indexes
GET _cat/indices

# Insert new document
# POST -> autogenerate id
POST cimri_product_index/_doc
{
  "title": "Apple iPhone 11",
  "category_id": 130
}

# PUT -> specific id
PUT cimri_product_index/_doc/1234
{
  "title": "Apple iPhone 12",
  "category_id": 130
}

# Get document
GET cimri_product_index/_doc/1234

# Override
PUT cimri_product_index/_doc/1234
{
  "title": "Apple iPhone 12 Beyaz",
  "category_id": 130
}

# Override istemiyorsak -> Create endpoint
PUT cimri_product_index/_create/1234
{
  "title": "Apple iPhone 12 Siyah",
  "category_id": 130
}

# Update
POST cimri_product_index/_update/1234
{
  "doc": {
    "title": "Apple iPhone 11 Beyaz 64GB"
  }
}

# Delete
DELETE cimri_product_index/_doc/1234
```
