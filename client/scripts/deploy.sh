#/bin/bash
#upload files
docker run --rm -it -v ~/.aws:/root/.aws -v $(pwd)/dist/fuse:/aws amazon/aws-cli --profile aws-fluxy-s3 s3 cp . s3://fluxy-production --recursive --include "*" --acl public-read --cache-control public,max-age=31536000,no-transform
docker run --rm -it -v ~/.aws:/root/.aws amazon/aws-cli --profile aws-fluxy-cloudfront cloudfront create-invalidation --distribution-id E378K7IAQVS246 --paths "/index.html"