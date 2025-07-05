import plotly.graph_objects as go
import json

# Data for the timeline
data = {
  "timeline": [
    {"phase": "Phase 1", "month": 1, "activity": "GoLang Team Training", "start": 0, "duration": 1},
    {"phase": "Phase 2", "month": 2, "activity": "Core System Development", "start": 1, "duration": 1}, 
    {"phase": "Phase 3", "month": 3, "activity": "Fyne GUI Development", "start": 2, "duration": 1},
    {"phase": "Phase 4", "month": 4, "activity": "Testing and Training", "start": 3, "duration": 1},
    {"phase": "Phase 5", "month": 5, "activity": "Deployment and Support", "start": 4, "duration": 1}
  ],
  "project_duration": 5
}

# Brand colors in order
colors = ['#1FB8CD', '#FFC185', '#ECEBD5', '#5D878F', '#D2BA4C']

# Create figure
fig = go.Figure()

# Add bars for each phase
for i, phase_data in enumerate(data['timeline']):
    # Truncate activity name to 15 characters as per instructions
    activity_short = phase_data['activity'][:15]
    fig.add_trace(go.Bar(
        y=[f"{phase_data['phase']}: {activity_short}"],
        x=[phase_data['duration']],
        base=[phase_data['start']],
        orientation='h',
        marker_color=colors[i],
        name=phase_data['phase'],
        hovertemplate=f"<b>{phase_data['phase']}</b><br>{phase_data['activity']}<br>Month {phase_data['month']}<extra></extra>",
        showlegend=False
    ))

# Update layout
fig.update_layout(
    title="Parish Bookstore System Timeline",
    xaxis_title="Month",
    yaxis_title="Project Phase",
    barmode='overlay'
)

# Update x-axis to show months
fig.update_xaxes(
    tickmode='linear',
    tick0=0.5,
    dtick=1,
    tickvals=[0.5, 1.5, 2.5, 3.5, 4.5],
    ticktext=['Month 1', 'Month 2', 'Month 3', 'Month 4', 'Month 5'],
    range=[0, 5]
)

# Update y-axis
fig.update_yaxes(
    autorange='reversed'  # Show Phase 1 at top
)

# Save the chart
fig.write_image("parish_bookstore_timeline.png")