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

## Part 5

```
# dynamic mapping
POST temp_index/_doc
{
  "name": "Pineapple",
  "botanical_name": "Ananas comosus",
  "produce_type": "Fruit",
  "country_of_origin": "New Zealand",
  "date_purchased": "2020-06-02T12:15:35",
  "quantity": 200,
  "unit_price": 3.11,
  "description": "a large juicy tropical fruit consisting of aromatic edible yellow flesh surrounded by a tough segmented skin and topped with a tuft of stiff leaves.These pineapples are sourced from New Zealand.",
  "vendor_details": {
    "vendor": "Tropical Fruit Growers of New Zealand",
    "main_contact": "Hugh Rose",
    "vendor_location": "Whangarei, New Zealand",
    "preferred_vendor": true
  }
}

GET temp_index/_mapping

# Text --> full-text search (texy analysis)
# Keyword --> exact searches, aggs, sorting

# Custom Mapping
## Rules

### - If you do not define a mapping ahead of time, Elasticsearch dynamically creates the mapping for you.
### - If you do decide to define your own mapping, you can do so at index creation.
### - ONE mapping is defined per index. Once the index has been created, we can only add new fields to a mapping. We CANNOT change the mapping of an existing field.
### - If you must change the type of an existing field, you must create a new index with the desired mapping, then reindex all documents into the new index.

## same document
POST test_index/_doc
{
  "name": "Pineapple",
  "botanical_name": "Ananas comosus",
  "produce_type": "Fruit",
  "country_of_origin": "New Zealand",
  "date_purchased": "2020-06-02T12:15:35",
  "quantity": 200,
  "unit_price": 3.11,
  "description": "a large juicy tropical fruit consisting of aromatic edible yellow flesh surrounded by a tough segmented skin and topped with a tuft of stiff leaves.These pineapples are sourced from New Zealand.",
  "vendor_details": {
    "vendor": "Tropical Fruit Growers of New Zealand",
    "main_contact": "Hugh Rose",
    "vendor_location": "Whangarei, New Zealand",
    "preferred_vendor": true
  }
}

GET temp_index/_mapping

# Copy mapping and remove "temp_index"
# change description, name --> text only
# change product_type --> keyword only
# botanical_name, vendor_details --> kullanılmayacak

PUT optimized_index
{
  "mappings": {
    "properties": {
      "botanical_name": {
        "enabled": false
      },
      "country_of_origin": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword"
          }
        }
      },
      "date_purchased": {
        "type": "date"
      },
      "description": {
        "type": "text"
      },
      "name": {
        "type": "text"
      },
      "produce_type": {
        "type": "keyword"
      },
      "quantity": {
        "type": "long"
      },
      "unit_price": {
        "type": "float"
      },
      "vendor_details": {
        "enabled": false
      }
    }
  }
}

GET optimized_index/_mapping

POST optimized_index/_doc
{
  "name": "Pineapple",
  "botanical_name": "Ananas comosus",
  "produce_type": "Fruit",
  "country_of_origin": "New Zealand",
  "date_purchased": "2020-06-02T12:15:35",
  "quantity": 200,
  "unit_price": 3.11,
  "description": "a large juicy tropical fruit consisting of aromatic edible yellow flesh surrounded by a tough segmented skin and topped with a tuft of stiff leaves.These pineapples are sourced from New Zealand.",
  "vendor_details": {
    "vendor": "Tropical Fruit Growers of New Zealand",
    "main_contact": "Hugh Rose",
    "vendor_location": "Whangarei, New Zealand",
    "preferred_vendor": true
  }
}

## extra field: "organic"
POST optimized_index/_doc
{
  "name": "Mango",
  "botanical_name": "Harum Manis",
  "produce_type": "Fruit",
  "country_of_origin": "Indonesia",
  "organic": true,
  "date_purchased": "2020-05-02T07:15:35",
  "quantity": 500,
  "unit_price": 1.5,
  "description": "Mango Arumanis or Harum Manis is originated from East Java. Arumanis means harum dan manis or fragrant and sweet just like its taste. The ripe Mango Arumanis has dark green skin coated with thin grayish natural wax. The flesh is deep yellow, thick, and soft with little to no fiber. Mango Arumanis is best eaten when ripe.",
  "vendor_details": {
    "vendor": "Ayra Shezan Trading",
    "main_contact": "Suharto",
    "vendor_location": "Binjai, Indonesia",
    "preferred_vendor": true
  }
}

GET optimized_index/_mapping

# Reindex

# yeni bir index yaratiriz
PUT optimized_v2
{
  "mappings": {
    "properties": {
      "botanical_name": {
        "type": "text"
      },
      "country_of_origin": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "date_purchased": {
        "type": "date"
      },
      "description": {
        "type": "text"
      },
      "name": {
        "type": "text"
      },
      "organic": {
        "type": "boolean"
      },
      "produce_type": {
        "type": "keyword"
      },
      "quantity": {
        "type": "long"
      },
      "unit_price": {
        "type": "float"
      },
      "vendor_details": {
        "type": "object",
        "enabled": false
      }
    }
  }
}

GET optimized_v2/_mapping

POST _reindex
{
  "source": {
    "index": "optimized_index"
  },
  "dest": {
    "index": "optimized_v2"
  }
}


## runtime
PUT optimized_v2/_mapping
{
  "runtime": {
    "total": {
      "type": "double",
      "script": {
        "source": "emit(doc['unit_price'].value* doc['quantity'].value)"
      }
    }
  }
}

GET optimized_v2/_mapping

GET optimized_v2/_search
{
  "size": 0,
  "aggs": {
    "total_expense": {
      "sum": {
        "field": "total"
      }
    }
  }
}
```

#### mapping text vs keyword

![](https://user-images.githubusercontent.com/60980933/122120075-324d5380-cde7-11eb-9b4e-744dfa6d527d.png)
