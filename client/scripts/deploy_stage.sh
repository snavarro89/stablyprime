#/bin/bash
#upload files
docker run --rm -it -v ~/.aws:/root/.aws -v $(pwd)/dist/stablyprime:/aws amazon/aws-cli --profile aws-stably-s3 s3 cp . s3://stableprime-stage --recursive --include "*" --acl public-read --cache-control public,max-age=31536000,no-transform 
docker run --rm -it -v ~/.aws:/root/.aws amazon/aws-cli --profile aws-stably-cloudfront cloudfront create-invalidation --distribution-id PENDING --paths "/index.html"