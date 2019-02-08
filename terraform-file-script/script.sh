for VAR in <all-the-repos-names-to-create>
do
	cp repo-before-pushing.tf "repo-$VAR.tf"
    sed -i -- "s/@{REPO-NAME}/$VAR/g" "repo-$VAR.tf"
    UNDER=$(echo "$VAR" | sed -r 's/-/_/g')   
    sed -i -- "s/@{MODULE-NAME}/$UNDER/g" "repo-$VAR.tf"
done