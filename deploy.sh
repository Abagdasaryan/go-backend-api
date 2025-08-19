#!/bin/bash

echo "🚀 Deploying Go Backend API to Railway..."

# Check if Railway CLI is installed
if ! command -v railway &> /dev/null; then
    echo "❌ Railway CLI not found. Installing..."
    npm install -g @railway/cli
fi

# Check if user is logged in
if ! railway whoami &> /dev/null; then
    echo "🔐 Please login to Railway..."
    railway login
fi

# Deploy to Railway
echo "📦 Deploying..."
railway up

echo "✅ Deployment complete!"
echo "🌐 Your API should be available at the Railway URL"
echo "🔍 Check the Railway dashboard for the exact URL"
