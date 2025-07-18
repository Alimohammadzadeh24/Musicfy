#!/bin/bash

# Script to install Git hooks

# Create hooks directory if it doesn't exist
mkdir -p .git/hooks

# Create post-checkout hook
cat > .git/hooks/post-checkout << 'EOF'
#!/bin/bash

# This hook is called after checkout is done
# It automatically sets up the environment based on the branch

# Get the current branch name
BRANCH=$(git rev-parse --abbrev-ref HEAD)
echo "Switched to branch: $BRANCH"

# Run branch setup if the script exists
if [ -f "./scripts/branch-setup.sh" ]; then
  echo "Setting up environment for branch $BRANCH..."
  ./scripts/branch-setup.sh
fi
EOF

# Make the hook executable
chmod +x .git/hooks/post-checkout

echo "Git hooks installed successfully"
echo "The environment will be automatically set up when switching branches" 