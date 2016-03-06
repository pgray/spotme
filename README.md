# spotme
Small go command line tool to request the cheapest spot instance available

The tool expects the following environment variables:
- AWS_DEFAULT_REGION
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY

usage:

-c, --config <file name>
-a, --ami <ami id>
-k, --key <key name>
-s, --security-group <security group id>
-i, --instance-types <list of instance types>
-t, --test
