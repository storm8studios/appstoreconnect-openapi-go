
generate:
	rm -rf generated
	/opt/homebrew/bin/openapi-generator generate \
	  -i https://raw.githubusercontent.com/EvanBacon/App-Store-Connect-OpenAPI-Spec/main/specs/latest.json \
	  -g go \
	  -c config/config.yml \
	  -o generated \
	  --skip-validate-spec \
	  --git-repo-id appstoreconnect-openapi-go/generated \
	  --git-user-id storm8studios \
	

