# Kibana Example Queries

---

# Part 1

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

---

# Part 2

```

# dataset --> https://www.kaggle.com/datasets/rmisra/news-category-dataset?resource=download

# Retrieve all documents from an index (sample of 10 by default)
GET news_headlines/_search

# Get total number
GET news_headlines/_search
{
  "track_total_hits": true
}

# Queries vs Aggregations
# Queries --> istenilen kriterlerdeki dökümanları getir

# Range
GET news_headlines/_search
{
  "query": {
    "range": {
      "date": {
        "gte": "2015-06-20",
        "lte": "2015-09-22"
      }
    }
  }
}

# Aggregations --> analiz et ve sonuçları topla
GET news_headlines/_search
{
  "aggs": {
    "by_category": {
      "terms": {
        "field": "category",
        "size": 100
      }
    }
  }
}

# Query + Aggregation
GET news_headlines/_search
{
  "query": {
    "match": {
      "category": "ENTERTAINMENT"
    }
  },
  "aggs": {
    "popular_in_entertainment": {
      "significant_text": {
        "field": "headline"
      }
    }
  }
}

# Increasing Consistency
GET news_headlines/_search
{
  "query": {
    "match": {
      "headline": {
        "query": "Khloe Kardashian Kendall Jenner"
      }
    }
  }
}
# match --> OR logic

GET news_headlines/_search
{
  "query": {
    "match": {
      "headline": {
        "query": "Khloe Kardashian Kendall Jenner",
        "operator": "and"
      }
    }
  }
}

GET news_headlines/_search
{
  "query": {
    "match": {
      "headline": {
        "query": "Khloe Kardashian Kendall Jenner",
        "minimum_should_match": 3
      }
    }
  }
}
```
