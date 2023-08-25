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

# Part 3

```
# Precision --> I want all the retrieved results to be a perfect match even if returns less document
# Recall --> I want to retrieve more results even if they not be a perfect match

# match query --> OR

# match --> alakasız sonuçlar da geliyor
GET news_headlines/_search
{
  "query": {
    "match": {
      "headline": {
        "query": "Shape of you"
      }
    }
  }
}

# match_phrase
GET news_headlines/_search
{
  "query": {
    "match_phrase": {
      "headline": {
        "query": "Shape of you"
      }
    }
  }
}

# multiple fields
# açıklamada geçiyor ama headline'ın alakası az sonuçlar önce geliyor
GET news_headlines/_search
{
  "query": {
    "multi_match": {
        "query": "Michelle Obama",
        "fields": [
          "headline",
          "short_description",
          "authors"
        ]
    }
  }
}

# boost
GET news_headlines/_search
{
  "query": {
    "multi_match": {
        "query": "Michelle Obama",
        "fields": [
          "headline^2",
          "short_description",
          "authors"
        ]
    }
  }
}

# multi_match --> OR şeklinde çalışıyor. Yine alakasız sonuçlar geliyor:
GET news_headlines/_search
{
  "query": {
    "multi_match": {
        "query": "party planning",
        "fields": [
          "headline^2",
          "short_description"
        ]
    }
  }
}

# multi_match + match_phrase
GET news_headlines/_search
{
  "query": {
    "multi_match": {
        "query": "party planning",
        "fields": [
          "headline^2",
          "short_description"
        ],
        "type": "phrase"
    }
  }
}

## Combined Queries --> Bool Query

# Headline'da "Michelle Obama" içeren "POLITICS" kategorisinden "2016" öncesi haberleri getir???

# Bool Query
# Must
GET news_headlines/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match_phrase": {
            "headline": "Michelle Obama"
          }
        },
        {
          "match": {
            "category": "POLITICS"
          }
        }
      ]
    }
  }
}

# Must Not
GET news_headlines/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match_phrase": {
            "headline": "Michelle Obama"
          }
        }
      ],
      "must_not": [
        {
          "match": {
            "category": "WEDDINGS"
          }
        }
      ]
    }
  }
}

# Should --> Nice to have (higher score)
GET news_headlines/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match_phrase": {
            "headline": "Michelle Obama"
          }
        }
      ],
      "should": [
        {
          "match_phrase": {
            "category": "BLACK VOICES"
          }
        }
      ]
    }
  }
}

# Filter
GET news_headlines/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match_phrase": {
            "headline": "Michelle Obama"
          }
        }
      ],
      "filter": {
        "range": {
          "date": {
            "gte": "2014-03-25",
            "lte": "2016-03-25"
          }
        }
      }
    }
  }
}
```

# Part 4

```
# Dataset --> https://www.kaggle.com/datasets/carrie1/ecommerce-data

GET ecommerce_data/_search

# Aggregations

GET ecommerce_data/_search
{
  "aggs": {
    "sum_unit_price": {
      "sum": {
        "field": "UnitPrice"
      }
    }
  }
}

# Hits gelmesin istiyorsak
GET ecommerce_data/_search
{
  "size": 0,
  "aggs": {
    "sum_unit_price": {
      "sum": {
        "field": "UnitPrice"
      }
    }
  }
}

GET ecommerce_data/_search
{
  "size": 0,
  "aggs": {
    "max_unit_price": {
      "max": {
        "field": "UnitPrice"
      }
    }
  }
}

GET ecommerce_data/_search
{
  "size": 0,
  "aggs": {
    "avg_unit_price": {
      "avg": {
        "field": "UnitPrice"
      }
    }
  }
}

# Cardinality
GET ecommerce_data/_search
{
  "size": 0,
  "aggs": {
    "num_of_unique_customers": {
      "cardinality": {
        "field": "CustomerID"
      }
    }
  }
}

# Query + Aggs
GET ecommerce_data/_search
{
  "size": 0,
  "query": {
    "match": {
      "Country": "Germany"
    }
  },
  "aggs": {
    "germany_avg_unit_price": {
      "avg": {
        "field": "UnitPrice"
      }
    }
  }
}

GET ecommerce_data/_search
{
  "size": 0,
  "aggs": {
    "top_5_customer": {
      "terms": {
        "field": "CustomerID",
        "size": 5
      }
    }
  }
}

# Order
GET ecommerce_data/_search
{
  "size": 0,
  "aggs": {
    "top_5_customer": {
      "terms": {
        "field": "CustomerID",
        "size": 5,
        "order": {
          "_count": "asc"
        }
      }
    }
  }
}

# Script
GET ecommerce_data/_search
{
  "size": 0,
  "aggs": {
    "daily_revenue": {
      "sum": {
        "script": {
          "source": "doc['UnitPrice'].value * doc['Quantity'].value"
        }
      }
    }
  }
}
```
