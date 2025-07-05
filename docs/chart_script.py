import plotly.graph_objects as go
import plotly.express as px

# Data for the system architecture
layers_data = [
    {
        "layer": "User Interface Layer",
        "components": ["Staff Interface (Fyne GUI)", "Public Kiosk Interface (Fyne GUI)", "Parish Council Reports Interface"],
        "color": "#2980b9"
    },
    {
        "layer": "Application Layer", 
        "components": ["Go Application Logic", "Inventory Management Service", "Reporting Service", "User Authentication Service"],
        "color": "#27ae60"
    },
    {
        "layer": "Data Layer",
        "components": ["Books Database", "Transactions Database", "Users Database", "Configuration Database"],
        "color": "#e67e22"
    }
]

# Prepare data for plotting
x_positions = []
y_positions = []
colors = []
hover_texts = []
component_names = []

# Create positions for each component
layer_y_positions = [3, 2, 1]  # Top to bottom
for i, layer in enumerate(layers_data):
    y_pos = layer_y_positions[i]
    components = layer["components"]
    layer_color = layer["color"]
    
    # Space components evenly across x-axis
    num_components = len(components)
    x_spacing = 1.0 / (num_components + 1)
    
    for j, component in enumerate(components):
        x_pos = (j + 1) * x_spacing
        x_positions.append(x_pos)
        y_positions.append(y_pos)
        colors.append(layer_color)
        
        # Abbreviate component names for display (15 char limit)
        short_name = component[:15]
        if len(component) > 15:
            short_name = component[:12] + "..."
        component_names.append(short_name)
        hover_texts.append(f"Layer: {layer['layer']}<br>Component: {component}")

# Create the scatter plot
fig = go.Figure()

# Add components as scatter points
fig.add_trace(go.Scatter(
    x=x_positions,
    y=y_positions,
    mode='markers+text',
    marker=dict(
        size=60,
        color=colors,
        symbol='square',
        line=dict(width=2, color='white')
    ),
    text=component_names,
    textposition="middle center",
    textfont=dict(size=10, color='white'),
    hovertemplate='%{hovertext}<extra></extra>',
    hovertext=hover_texts,
    showlegend=False
))

# Add layer labels on the right side
layer_names = ["UI Layer", "App Layer", "Data Layer"]
fig.add_trace(go.Scatter(
    x=[1.1, 1.1, 1.1],
    y=[3, 2, 1],
    mode='text',
    text=layer_names,
    textfont=dict(size=14, color='black'),
    showlegend=False,
    hoverinfo='skip'
))

# Add connecting arrows between layers
for i in range(len(layer_y_positions) - 1):
    fig.add_shape(
        type="line",
        x0=0.5, y0=layer_y_positions[i] - 0.3,
        x1=0.5, y1=layer_y_positions[i+1] + 0.3,
        line=dict(
            color="gray",
            width=3,
            dash="dot"
        )
    )
    
    # Add arrow head
    fig.add_shape(
        type="line",
        x0=0.45, y0=layer_y_positions[i+1] + 0.35,
        x1=0.5, y1=layer_y_positions[i+1] + 0.3,
        line=dict(color="gray", width=3)
    )
    fig.add_shape(
        type="line",
        x0=0.55, y0=layer_y_positions[i+1] + 0.35,
        x1=0.5, y1=layer_y_positions[i+1] + 0.3,
        line=dict(color="gray", width=3)
    )

# Update layout
fig.update_layout(
    title="Parish Bookstore System Architecture",
    xaxis=dict(
        showgrid=False,
        showticklabels=False,
        zeroline=False,
        range=[-0.1, 1.2]
    ),
    yaxis=dict(
        showgrid=False,
        showticklabels=False,
        zeroline=False,
        range=[0.5, 3.5]
    ),
    plot_bgcolor='white',
    paper_bgcolor='white'
)

# Save the chart
fig.write_image("parish_bookstore_architecture.png")