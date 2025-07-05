import pandas as pd
import json

# Create comprehensive project data for the parish council
project_data = {
    'Project Overview': {
        'Project Name': 'Parish Bookstore Inventory Management System',
        'Technology Stack': 'GoLang + Fyne GUI Framework',
        'Development Timeline': '5 months',
        'Target Users': 'Parish staff, volunteers, parishioners, council',
        'Primary Goal': 'Modernize inventory tracking and reporting'
    },
    
    'Technical Requirements': {
        'Programming Language': 'Go (Golang)',
        'GUI Framework': 'Fyne 2.x',
        'Database': 'SQLite/PostgreSQL',
        'Platform': 'Desktop (Windows, Mac, Linux)',
        'Hardware Requirements': 'Standard PC/laptop, barcode scanner (optional)'
    },
    
    'Financial Analysis': {
        'Development Cost Estimate': '$15,000 - $20,000',
        'Annual Operating Cost': '$1,200 - $1,800',
        'Annual Benefits': '$9,960',
        'ROI Timeline': '18-24 months',
        'Break-even Point': 'Month 20'
    },
    
    'Key Features': [
        'Real-time inventory tracking',
        'Automated stock level alerts',
        'Parish council financial reports',
        'Barcode scanning support',
        'User-friendly interface for volunteers',
        'Price and availability lookup for parishioners',
        'Sales transaction recording',
        'Monthly/quarterly reporting',
        'Backup and data recovery',
        'Multi-user access control'
    ],
    
    'Benefits for Stakeholders': {
        'Parish Council': [
            'Accurate financial reporting',
            'Real-time inventory valuation',
            'Clear profit/loss tracking',
            'Automated monthly reports',
            'Better budget planning'
        ],
        'Parishioners': [
            'Easy price checking',
            'Stock availability information',
            'Faster service',
            'Better book selection',
            'Professional experience'
        ],
        'Staff/Volunteers': [
            '30% reduction in manual work',
            'Simplified checkout process',
            'Automatic inventory updates',
            'Less paperwork',
            'Training provided'
        ]
    },
    
    'Implementation Timeline': [
        {'Phase': 1, 'Duration': 'Month 1', 'Activity': 'Team training on GoLang', 'Deliverable': 'Trained development team'},
        {'Phase': 2, 'Duration': 'Month 2', 'Activity': 'Core system development', 'Deliverable': 'Database and backend logic'},
        {'Phase': 3, 'Duration': 'Month 3', 'Activity': 'Fyne GUI development', 'Deliverable': 'User interface and workflows'},
        {'Phase': 4, 'Duration': 'Month 4', 'Activity': 'Testing and user training', 'Deliverable': 'Tested system and trained users'},
        {'Phase': 5, 'Duration': 'Month 5', 'Activity': 'Deployment and support', 'Deliverable': 'Live system with documentation'}
    ],
    
    'Risk Assessment': {
        'Technical Risks': [
            'Learning curve for GoLang (Mitigation: Training and documentation)',
            'Fyne framework limitations (Mitigation: Prototype early)',
            'Data migration challenges (Mitigation: Careful planning)'
        ],
        'Operational Risks': [
            'User adoption resistance (Mitigation: Training and support)',
            'System downtime (Mitigation: Backup procedures)',
            'Data loss (Mitigation: Regular backups)'
        ],
        'Financial Risks': [
            'Budget overrun (Mitigation: Fixed-price contract)',
            'Delayed ROI (Mitigation: Phased implementation)'
        ]
    }
}

# Create detailed breakdown tables
# 1. Financial Benefits Breakdown
benefits_df = pd.DataFrame([
    {'Benefit Category': 'Time Savings (Volunteer Hours)', 'Annual Value': '$3,600', 'Description': '20-30% reduction in manual inventory tasks'},
    {'Benefit Category': 'Improved Inventory Turnover', 'Annual Value': '$2,400', 'Description': 'Better stock management reduces dead inventory'},
    {'Benefit Category': 'Reduced Stock Losses', 'Annual Value': '$1,800', 'Description': 'Less theft, damage, and obsolescence'},
    {'Benefit Category': 'Better Cash Flow Management', 'Annual Value': '$1,200', 'Description': 'Optimized ordering and reduced overstock'},
    {'Benefit Category': 'Automated Reporting', 'Annual Value': '$960', 'Description': 'Reduced administrative time for reports'}
])

# 2. GoLang Learning Path
golang_learning_df = pd.DataFrame([
    {'Week': 1, 'Topic': 'Go Basics & Syntax', 'Hours': 10, 'Resources': 'Official Go Tour, basic tutorials'},
    {'Week': 2, 'Topic': 'Data Structures & Functions', 'Hours': 12, 'Resources': 'Go by Example, practice exercises'},
    {'Week': 3, 'Topic': 'Database Integration', 'Hours': 15, 'Resources': 'SQL drivers, database/sql package'},
    {'Week': 4, 'Topic': 'Fyne GUI Development', 'Hours': 20, 'Resources': 'Fyne documentation, sample projects'},
    {'Week': 5, 'Topic': 'Project Implementation', 'Hours': 25, 'Resources': 'Hands-on development work'},
    {'Week': 6, 'Topic': 'Testing & Debugging', 'Hours': 15, 'Resources': 'Go testing package, debugging tools'}
])

# 3. System Requirements
requirements_df = pd.DataFrame([
    {'Component': 'Hardware', 'Specification': 'Standard PC/Laptop', 'Cost': '$0 (existing)', 'Notes': 'Windows 10+, macOS 10.14+, or Linux'},
    {'Component': 'Barcode Scanner', 'Specification': 'USB Handheld Scanner', 'Cost': '$50-100', 'Notes': 'Optional for faster inventory updates'},
    {'Component': 'Software License', 'Specification': 'Go + Fyne (Open Source)', 'Cost': '$0', 'Notes': 'No licensing fees'},
    {'Component': 'Database', 'Specification': 'SQLite or PostgreSQL', 'Cost': '$0', 'Notes': 'Free, reliable options'},
    {'Component': 'Development Tools', 'Specification': 'VS Code or GoLand IDE', 'Cost': '$0-200/year', 'Notes': 'Professional development environment'}
])

# 4. Comparison with Current System
comparison_df = pd.DataFrame([
    {'Aspect': 'Inventory Tracking', 'Current System': 'Manual spreadsheets', 'New System': 'Automated real-time tracking', 'Improvement': 'Accurate, always up-to-date'},
    {'Aspect': 'Stock Counts', 'Current System': 'Quarterly manual counts', 'New System': 'Continuous automated updates', 'Improvement': '95%+ accuracy vs ~70%'},
    {'Aspect': 'Council Reports', 'Current System': 'Manual monthly reports', 'New System': 'Automated generated reports', 'Improvement': 'Save 8-10 hours/month'},
    {'Aspect': 'Price Lookup', 'Current System': 'Check physical books/lists', 'New System': 'Instant computer lookup', 'Improvement': 'Faster customer service'},
    {'Aspect': 'Sales Recording', 'Current System': 'Cash box + manual log', 'New System': 'Integrated POS system', 'Improvement': 'Automatic transaction recording'},
    {'Aspect': 'Stock Alerts', 'Current System': 'Visual inspection', 'New System': 'Automated low-stock alerts', 'Improvement': 'Never run out of popular items'}
])

# Save all data to CSV files
benefits_df.to_csv('parish_bookstore_benefits.csv', index=False)
golang_learning_df.to_csv('golang_learning_path.csv', index=False)
requirements_df.to_csv('system_requirements.csv', index=False) 
comparison_df.to_csv('current_vs_new_system.csv', index=False)

# Create a summary report
with open('parish_bookstore_project_summary.txt', 'w') as f:
    f.write("PARISH BOOKSTORE INVENTORY MANAGEMENT SYSTEM\n")
    f.write("="*50 + "\n\n")
    
    f.write("PROJECT OVERVIEW\n")
    f.write("-" * 20 + "\n")
    for key, value in project_data['Project Overview'].items():
        f.write(f"{key}: {value}\n")
    f.write("\n")
    
    f.write("FINANCIAL ANALYSIS\n")
    f.write("-" * 20 + "\n")
    for key, value in project_data['Financial Analysis'].items():
        f.write(f"{key}: {value}\n")
    f.write("\n")
    
    f.write("KEY FEATURES\n")
    f.write("-" * 20 + "\n")
    for i, feature in enumerate(project_data['Key Features'], 1):
        f.write(f"{i}. {feature}\n")
    f.write("\n")
    
    f.write("IMPLEMENTATION PHASES\n")
    f.write("-" * 20 + "\n")
    for phase in project_data['Implementation Timeline']:
        f.write(f"Phase {phase['Phase']} ({phase['Duration']}): {phase['Activity']}\n")
        f.write(f"  Deliverable: {phase['Deliverable']}\n")
    f.write("\n")

print("Parish Bookstore Project Files Created:")
print("1. parish_bookstore_benefits.csv - Financial benefits breakdown")
print("2. golang_learning_path.csv - GoLang training schedule")
print("3. system_requirements.csv - Technical requirements and costs")
print("4. current_vs_new_system.csv - Comparison analysis")
print("5. parish_bookstore_project_summary.txt - Executive summary")

# Display sample data
print("\nSample Financial Benefits:")
print(benefits_df.to_string(index=False))

print("\nSample Learning Path:")
print(golang_learning_df.head(3).to_string(index=False))