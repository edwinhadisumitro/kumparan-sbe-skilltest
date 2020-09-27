# kumparan-sbe-skilltest
Kumparan Senior Backend Engineer Skilltest

Deployment : 
- Execute run.sh, kumparan-news-api and kumparan-news-listener services will not start because they need to establish a connection to ElasticSearch, which take some time to be ready
- Wait for ElasticSearch container to be ready for query.
- Execute docker-compose up -d to start kumparan-news-api and kumparan-news-listener services

