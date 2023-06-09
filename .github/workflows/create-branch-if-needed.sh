if [ $(git status --porcelain | wc -l) -eq "0" ]; then
  echo "🟢 Git repo is clean."
else
  echo "🔴 Git repo modified."
  timestamp=$(date +%s)
  branchName="regen/clients-${timestamp}"
  git checkout -b $branchName
  git add --all
  git commit -a -m "Regenerated API clients with Kiota"
fi
