if [ $(git status --porcelain | wc -l) -eq "0" ]; then
  echo "🟢 Git repo is clean."
else
  echo "🔴 Git repo modified."
  timestamp=$(date +%s)
  branchName="regen/clients-${timestamp}"
  git checkout -b $branchName
  git add --all
  git commit -a -m "Regenerated API clients with Kiota"
  git push --set-upstream origin $branchName
  gh issue create --title "Validate regenerated Kiota clients in get started" --body "Changes in branch $branchName" --repo microsoft/kiota-samples --assignee jasonjoh
fi
