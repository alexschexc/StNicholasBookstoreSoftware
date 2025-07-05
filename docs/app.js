// Slide Presentation JavaScript
class SlidePresentation {
    constructor() {
        this.currentSlide = 1;
        this.totalSlides = 10;
        this.slides = document.querySelectorAll('.slide');
        this.prevBtn = document.getElementById('prev-btn');
        this.nextBtn = document.getElementById('next-btn');
        this.currentSlideSpan = document.getElementById('current-slide');
        this.totalSlidesSpan = document.getElementById('total-slides');
        
        this.init();
    }
    
    init() {
        // Set initial values
        this.updateSlideCounter();
        this.updateButtonStates();
        
        // Bind event listeners
        this.prevBtn.addEventListener('click', () => this.previousSlide());
        this.nextBtn.addEventListener('click', () => this.nextSlide());
        
        // Add keyboard navigation
        document.addEventListener('keydown', (e) => this.handleKeydown(e));
        
        // Ensure first slide is visible
        this.showSlide(1);
    }
    
    showSlide(slideNumber) {
        // Hide all slides
        this.slides.forEach(slide => {
            slide.classList.remove('active');
        });
        
        // Show current slide
        const targetSlide = document.getElementById(`slide-${slideNumber}`);
        if (targetSlide) {
            targetSlide.classList.add('active');
        }
        
        this.currentSlide = slideNumber;
        this.updateSlideCounter();
        this.updateButtonStates();
    }
    
    nextSlide() {
        if (this.currentSlide < this.totalSlides) {
            this.showSlide(this.currentSlide + 1);
        }
    }
    
    previousSlide() {
        if (this.currentSlide > 1) {
            this.showSlide(this.currentSlide - 1);
        }
    }
    
    updateSlideCounter() {
        this.currentSlideSpan.textContent = this.currentSlide;
        this.totalSlidesSpan.textContent = this.totalSlides;
    }
    
    updateButtonStates() {
        // Update Previous button
        if (this.currentSlide <= 1) {
            this.prevBtn.disabled = true;
            this.prevBtn.classList.add('disabled');
        } else {
            this.prevBtn.disabled = false;
            this.prevBtn.classList.remove('disabled');
        }
        
        // Update Next button
        if (this.currentSlide >= this.totalSlides) {
            this.nextBtn.disabled = true;
            this.nextBtn.classList.add('disabled');
        } else {
            this.nextBtn.disabled = false;
            this.nextBtn.classList.remove('disabled');
        }
    }
    
    handleKeydown(event) {
        // Handle keyboard navigation
        switch(event.key) {
            case 'ArrowRight':
            case ' ': // Spacebar
                event.preventDefault();
                this.nextSlide();
                break;
            case 'ArrowLeft':
                event.preventDefault();
                this.previousSlide();
                break;
            case 'Home':
                event.preventDefault();
                this.showSlide(1);
                break;
            case 'End':
                event.preventDefault();
                this.showSlide(this.totalSlides);
                break;
            case 'Escape':
                // Optional: Add fullscreen exit or other functionality
                break;
        }
    }
    
    // Method to go to a specific slide (for future enhancements)
    goToSlide(slideNumber) {
        if (slideNumber >= 1 && slideNumber <= this.totalSlides) {
            this.showSlide(slideNumber);
        }
    }
}

// Additional utility functions for enhanced presentation features
class PresentationUtils {
    static formatSlideContent() {
        // Add any content formatting logic here
        // For example, auto-numbering lists, formatting code blocks, etc.
    }
    
    static addAccessibilityFeatures() {
        // Add ARIA labels and roles for better accessibility
        const slides = document.querySelectorAll('.slide');
        slides.forEach((slide, index) => {
            slide.setAttribute('role', 'tabpanel');
            slide.setAttribute('aria-label', `Slide ${index + 1} of ${slides.length}`);
        });
        
        // Add aria-live region for slide announcements
        const slideAnnouncer = document.createElement('div');
        slideAnnouncer.setAttribute('aria-live', 'polite');
        slideAnnouncer.setAttribute('aria-atomic', 'true');
        slideAnnouncer.className = 'sr-only';
        slideAnnouncer.id = 'slide-announcer';
        document.body.appendChild(slideAnnouncer);
    }
    
    static announceSlideChange(slideNumber, totalSlides) {
        const announcer = document.getElementById('slide-announcer');
        if (announcer) {
            announcer.textContent = `Slide ${slideNumber} of ${totalSlides}`;
        }
    }
    
    static addSlideTransitionEffects() {
        // Add custom transition effects for slides
        const slides = document.querySelectorAll('.slide');
        slides.forEach(slide => {
            slide.style.transition = 'all 0.5s cubic-bezier(0.16, 1, 0.3, 1)';
        });
    }
    
    static handlePrintMode() {
        // Add print-specific styling
        const printStyles = `
            @media print {
                .slide {
                    position: static !important;
                    opacity: 1 !important;
                    transform: none !important;
                    page-break-after: always;
                    height: auto !important;
                    padding: 20px !important;
                }
                .nav-controls {
                    display: none !important;
                }
                .slide-content {
                    height: auto !important;
                }
            }
        `;
        
        const styleSheet = document.createElement('style');
        styleSheet.textContent = printStyles;
        document.head.appendChild(styleSheet);
    }
}

// Enhanced slide presentation with additional features
class EnhancedSlidePresentation extends SlidePresentation {
    constructor() {
        super();
        this.autoAdvanceInterval = null;
        this.isAutoAdvancing = false;
        this.slideHistory = [];
        
        this.addEnhancedFeatures();
    }
    
    addEnhancedFeatures() {
        // Add progress indicator
        this.addProgressIndicator();
        
        // Add slide overview/thumbnail navigation (optional)
        this.addSlideOverview();
        
        // Add presentation timer
        this.addPresentationTimer();
    }
    
    addProgressIndicator() {
        const progressBar = document.createElement('div');
        progressBar.className = 'progress-bar';
        progressBar.innerHTML = '<div class="progress-fill"></div>';
        
        const progressStyles = `
            .progress-bar {
                position: fixed;
                top: 0;
                left: 0;
                width: 100%;
                height: 4px;
                background: rgba(0, 0, 0, 0.1);
                z-index: 1000;
            }
            .progress-fill {
                height: 100%;
                background: var(--color-presentation-primary);
                transition: width 0.3s ease;
                width: ${(this.currentSlide / this.totalSlides) * 100}%;
            }
        `;
        
        const styleSheet = document.createElement('style');
        styleSheet.textContent = progressStyles;
        document.head.appendChild(styleSheet);
        document.body.appendChild(progressBar);
    }
    
    updateProgressBar() {
        const progressFill = document.querySelector('.progress-fill');
        if (progressFill) {
            progressFill.style.width = `${(this.currentSlide / this.totalSlides) * 100}%`;
        }
    }
    
    addSlideOverview() {
        // Create a slide overview toggle button
        const overviewBtn = document.createElement('button');
        overviewBtn.className = 'btn btn--outline overview-btn';
        overviewBtn.textContent = 'Overview';
        overviewBtn.addEventListener('click', () => this.toggleSlideOverview());
        
        // Add to navigation controls
        const navControls = document.querySelector('.nav-controls');
        navControls.insertBefore(overviewBtn, navControls.querySelector('.slide-counter'));
    }
    
    toggleSlideOverview() {
        // Implementation for slide overview/grid view
        console.log('Slide overview toggled');
    }
    
    addPresentationTimer() {
        const timer = document.createElement('div');
        timer.className = 'presentation-timer';
        timer.textContent = '00:00';
        
        const timerStyles = `
            .presentation-timer {
                position: fixed;
                top: 20px;
                right: 20px;
                background: var(--color-presentation-card);
                padding: 8px 12px;
                border-radius: 4px;
                font-family: var(--font-family-mono);
                font-size: var(--font-size-sm);
                color: var(--color-text-secondary);
                border: 1px solid var(--color-border);
                box-shadow: var(--shadow-sm);
            }
        `;
        
        const styleSheet = document.createElement('style');
        styleSheet.textContent = timerStyles;
        document.head.appendChild(styleSheet);
        document.body.appendChild(timer);
        
        this.startPresentationTimer();
    }
    
    startPresentationTimer() {
        const startTime = Date.now();
        const timer = document.querySelector('.presentation-timer');
        
        setInterval(() => {
            const elapsed = Date.now() - startTime;
            const minutes = Math.floor(elapsed / 60000);
            const seconds = Math.floor((elapsed % 60000) / 1000);
            timer.textContent = `${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
        }, 1000);
    }
    
    showSlide(slideNumber) {
        super.showSlide(slideNumber);
        this.updateProgressBar();
        this.addToSlideHistory(slideNumber);
        
        // Announce slide change for accessibility
        PresentationUtils.announceSlideChange(slideNumber, this.totalSlides);
    }
    
    addToSlideHistory(slideNumber) {
        this.slideHistory.push({
            slideNumber: slideNumber,
            timestamp: Date.now()
        });
    }
    
    // Auto-advance functionality
    startAutoAdvance(intervalMs = 30000) {
        if (this.isAutoAdvancing) return;
        
        this.isAutoAdvancing = true;
        this.autoAdvanceInterval = setInterval(() => {
            if (this.currentSlide < this.totalSlides) {
                this.nextSlide();
            } else {
                this.stopAutoAdvance();
            }
        }, intervalMs);
    }
    
    stopAutoAdvance() {
        if (this.autoAdvanceInterval) {
            clearInterval(this.autoAdvanceInterval);
            this.autoAdvanceInterval = null;
            this.isAutoAdvancing = false;
        }
    }
    
    // Method to export presentation data
    exportPresentationData() {
        return {
            currentSlide: this.currentSlide,
            totalSlides: this.totalSlides,
            slideHistory: this.slideHistory,
            sessionData: {
                startTime: this.slideHistory[0]?.timestamp || Date.now(),
                endTime: Date.now()
            }
        };
    }
}

// Initialize the presentation when the DOM is loaded
document.addEventListener('DOMContentLoaded', function() {
    // Add accessibility features
    PresentationUtils.addAccessibilityFeatures();
    PresentationUtils.addSlideTransitionEffects();
    PresentationUtils.handlePrintMode();
    
    // Initialize the presentation
    const presentation = new EnhancedSlidePresentation();
    
    // Make presentation instance globally available for debugging
    window.presentation = presentation;
    
    // Add some helpful console messages
    console.log('Parish Bookstore Presentation loaded successfully!');
    console.log('Keyboard shortcuts:');
    console.log('  • Arrow keys: Navigate slides');
    console.log('  • Spacebar: Next slide');
    console.log('  • Home: Go to first slide');
    console.log('  • End: Go to last slide');
    
    // Handle window resize for responsive design
    window.addEventListener('resize', function() {
        // Recalculate slide dimensions if needed
        const slides = document.querySelectorAll('.slide');
        slides.forEach(slide => {
            // Trigger reflow for any layout adjustments
            slide.style.display = 'none';
            slide.offsetHeight; // Trigger reflow
            slide.style.display = '';
        });
    });
    
    // Handle visibility change (for presentation mode)
    document.addEventListener('visibilitychange', function() {
        if (document.hidden) {
            // Pause any auto-advance if the tab is not visible
            if (presentation.isAutoAdvancing) {
                presentation.stopAutoAdvance();
            }
        }
    });
});

// Export for potential module usage
if (typeof module !== 'undefined' && module.exports) {
    module.exports = { SlidePresentation, EnhancedSlidePresentation, PresentationUtils };
}