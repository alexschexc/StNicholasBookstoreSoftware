import plotly.graph_objects as go
import json

# Data from the provided JSON
data = {
  "benefits": [
    {"category": "Time Savings", "value": 3600, "description": "Volunteer hours reduced"},
    {"category": "Improved Inventory Turnover", "value": 2400, "description": "Better stock management"}, 
    {"category": "Reduced Stock Losses", "value": 1800, "description": "Less theft and damage"},
    {"category": "Better Cash Flow", "value": 1200, "description": "Optimized ordering"},
    {"category": "Automated Reporting", "value": 960, "description": "Staff time savings"}
  ],
  "total_annual_benefit": 9960
}

# Extract data for the chart
categories = []
values = []
for benefit in data["benefits"]:
    # Shorten category names to meet 15 character limit
    if benefit["category"] == "Time Savings":
        categories.append("Time Savings")
    elif benefit["category"] == "Improved Inventory Turnover":
        categories.append("Improved Invtry")
    elif benefit["category"] == "Reduced Stock Losses":
        categories.append("Reduced Losses")
    elif benefit["category"] == "Better Cash Flow":
        categories.append("Better Cash")
    elif benefit["category"] == "Automated Reporting":
        categories.append("Auto Reporting")
    
    values.append(benefit["value"])

# Brand colors as specified in instructions
colors = ['#1FB8CD', '#FFC185', '#ECEBD5', '#5D878F', '#D2BA4C']

# Create the bar chart
fig = go.Figure(data=[
    go.Bar(
        x=categories,
        y=values,
        marker_color=colors[:len(categories)],
        text=[f'${v/1000:.1f}k' for v in values],  # Format values with k for thousands
        textposition='outside',
        cliponaxis=False
    )
])

# Update layout with shortened title (under 40 characters)
fig.update_layout(
    title="Parish Inventory System Benefits",
    xaxis_title="Benefit Category",
    yaxis_title="Annual Value ($)",
    showlegend=False
)

# Format y-axis to show values in thousands
fig.update_yaxes(
    tickformat='$,.0f'
)

# Save the chart
fig.write_image("parish_inventory_benefits.png")