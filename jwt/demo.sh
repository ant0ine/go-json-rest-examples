#!/usr/bin/env bash

################################################################################
# Helper functions
#   These are just to clean up the demo code.
################################################################################


# get_jwt user pass
#   Attemptes to obtain a JWT token from the current running service
function get_jwt() {
  curl --silent -d "{\"username\": \"$1\", \"password\": \"$2\"}" -H "Content-Type:application/json" http://localhost:8080/api/login | grep token | cut -d: -f2 | sed -e 's/^ *//g' -e 's/ *$//g' | tr -d '"'
}
# api_get JWT api_endpoint
#   Wrapper to query running service with given token
function api_get() {
  curl -H "Authorization:Bearer $1" http://localhost:8080/api/$2
}
# pretty_get LOG_PREIX API_ENDPOINT JWT
#   Helper to output stuff to the console in a pretty/uniform way
function pretty_get() {
  echo -e "$1: /api/$2"
  api_get $3 $2
  echo ""
}


################################################################################
# Obtain JWT Tokens
################################################################################
cat <<-EOF
--------------------------------------------------------------------------------
- /api/login / TOKENS
--------------------------------------------------------------------------------
EOF
JWT_ADMIN=$(get_jwt admin admin)
JWT_USER=$(get_jwt user pass)
# Here is an admin token that expires in 2099 (http://jwt.io), but has a bad signature.
JWT_ADMIN_BADSIG="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQwODk2NTc2MDAsImlkIjoiYWRtaW4iLCJpc0FkbWluIjp0cnVlLCJvcmlnX2lhdCI6MTQzMjEzMzM1OX0.BADSIG"
# Echo the obtained tokens back
echo "FAILED ATTEMPT:"; api_get user bad_pass
echo -e "\n\nADMIN TOKEN:\n$JWT_ADMIN\n\nUSER TOKEN:\n$JWT_USER"


################################################################################
# API Requests
################################################################################
### ENDPOINT: /api/auth_test
cat <<-EOF

--------------------------------------------------------------------------------
- /api/auth_test
--------------------------------------------------------------------------------
EOF
pretty_get "USER"       auth_test $JWT_USER
pretty_get "ADMIN"      auth_test $JWT_ADMIN
pretty_get "ADMIN BAD"  auth_test $JWT_ADMIN_BADSIG

### ENDPOINT: /api/admin_only
cat <<-EOF

--------------------------------------------------------------------------------
- /api/admin_only
--------------------------------------------------------------------------------
EOF
pretty_get "USER"       admin_only $JWT_USER
pretty_get "ADMIN"      admin_only $JWT_ADMIN
pretty_get "ADMIN BAD"  admin_only $JWT_ADMIN_BAD

### ENDPOINT: /api/refresh_token
cat <<-EOF

--------------------------------------------------------------------------------
- /api/refresh_token
--------------------------------------------------------------------------------
EOF
echo "Sleeping for 1 second so that token will change"
sleep 1
pretty_get "USER"       refresh_token $JWT_USER
