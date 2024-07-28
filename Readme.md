# BrewScylla
This is a coffee house backend for [the frontend app]().

## Design Choices
- Using **CassandraDB** as a proof of concept (PoC) to explore its capabilities and performance. 
- **Levenshtein** is used for search capabilities, allowing users to perform fast and flexible searches on product names, which is in-memory instead of using a specific service like ElasticSearch or Lucene.
- The web application is built with **Go and the Fiber framework** to handle requests.

## DB schema
Although CassandraDB is a NoSQL database, it operates as a wide-column store. In this model, the schema requires that columns be pre-defined. This means that, unlike some other NoSQL databases that offer flexible or dynamic schemas, the structure of each table in ScyllaDB is fixed and must be defined at creation time.

The schema used in this project is:

![db schema image](images/DB%20schema.png)

Some important considerations while designing wide-column NoSQL DBs
1. Filtering can only be done on Primary columns or Clustering columns.
2. Primary keys are always cumpolsary while querying, but clustering keys can be optional.

    1. If any clustering columns are used in a query, you must specify all previous clustering columns (from left to right) with equality comparisons (=) before using any inequality comparisons (<, >, <=, >=, IN).
    2. Clustering columns can be skipped from right to left. This means you can use fewer clustering columns than defined, but you must start from the leftmost clustering column and cannot skip columns in between.

Credits to DataStax Data Modelling [documentation](https://www.datastax.com/dev/modeling) for helping with data modelling best practices.

# TODOs
1. Handling UUIDs
2. User update, coupons. Handling search products.
3. Make retries work with gocql [link](https://stackoverflow.com/questions/76833860/how-do-i-make-the-cassandra-gocql-retry-policy-work) 
4. Admin APIs