---
layout: post
title: "How to setup a similar page"
date: 2025-09-03
author: "Paul Gondolf (with the help of Claude)"
excerpt: "A guide to setting up and customizing this Jekyll-based academic website template."
---

This guide explains how to set up and customize this Jekyll-based website template.

## Quick Start

### Prerequisites

- Ruby (2.7+)
- Bundler gem
- Git

### Installation

1. Fork or clone this repository
2. Navigate to the project directory
3. Install dependencies:

   ```bash
   bundle install
   ```

4. Start the development server:

   ```bash
   bundle exec jekyll serve
   ```

5. Visit `http://localhost:4000` to see your site

## Configuration

### Site Settings (_config.yml)

Update the main configuration file with your information:

```yaml
baseurl: "/websitereponame" # the subpath of your site (the repo name)
url: "https://yourusername.github.io" # the base hostname & protocol for your site
title: "Your Name"
email: "your.email@example.com"
description: >-
  Your professional description here.
url: "https://yourusername.github.io"
github_username: "your-github-username"
orcid_id: "0000-0000-0000-0000"
email_user: "youremailusername"
email_domain: "example.com"
```

Of course you should also update the cv.md, index.md, lecture_notes.md, research.md and talks.md with your own information.

### Additional Customizations

Don't forget to personalize these visual elements:

- **Profile Image**: Replace the default profile image in `/assets/images/profile.jpg` with your own professional photo
- **Favicon**: Create and add a favicon (the small icon in browser tabs) - you can use your initials, a professional headshot, or a symbol related to your field. Place favicon files in `/assets/images/` and add the appropriate `<link>` tags to your layout's `<head>` section

In the robots.txt you should change the link to the Sitemap and also set your privacy settings.

```txt
User-agent: *
Allow: /

Sitemap: https://yourusername.github.io/websitereponame/sitemap.xml
```

## Page Templates & Usage

### CV Page (cv.md)

The CV template supports structured data for education, teaching, and work experience:

```yaml
---
layout: cv
title: "Curriculum Vitae"
permalink: /cv/
last_updated: "August 2025"

education:
  - date: "Present"
    degree: "PhD in Mathematics"
    institution: "University Name"
    institution_url: "https://university.edu/"
    location: "City, Country"
    details:
      - label: "Expected Graduation"
        text: "Month Year"
      - label: "Thesis"
        text: "Your thesis title"
        file: "/path/to/thesis.pdf" # or 
        url: "https://url-to-your-thesis.com"
      - label: "Research Focus"
        text: "Your research area"

teaching:
  - title: "Course Title"
    period: "2022 - 2024"
    institution: "University Name"
    institution_url: "https://university.edu/"
    responsibilities:
      - "Responsibility 1"
      - "Responsibility 2"

work:
  - date: "2021-2022"
    position: "Job Title"
    company: "Company Name"
    company_url: "https://company.com/"
    location: "City, Country"
    description: "Brief job description"
    achievements:
      - "Achievement 1"
      - "Achievement 2"
---
```

### Research Page (research.md)

The research template organizes publications and preprints:

```yaml
---
layout: research
title: "Research"
permalink: /research/
description: "Brief description of your research"

published_papers:
  - title: "Paper Title"
    journal: "Journal Name"
    year: 2024
    authors:
      - name: "Coauthor Name"
      - name: "Your Name"
        highlight: true  # Highlights your name
    abstract: "Paper abstract here..."
    award:  # Optional
      title: "Award Name"
      url: "https://award-url.com"
      description: "Award description"
    links:
      - text: "arXiv:1234.5678"
        url: "https://arxiv.org/abs/1234.5678"
      - text: "DOI"
        url: "https://doi.org/..."

preprints:
  - title: "Preprint Title"
    status: "Submitted to Journal Name"
    year: 2024
    authors: # [same format as above]
    abstract: "Abstract..."
    links: # [same format as above]

research_interests:
  - area: "Research Area 1"
    description: "Description of this research area"
  - area: "Research Area 2"
    description: "Description of this research area"
---
```

### Talks Page (talks.md)

Create a talks page following this structure:

```yaml
---
layout: talks
title: "Talks"
permalink: /talks/
description: "Conference presentations and seminars"

talks:
  - title: "Talk Title"
    meta: # all optional 
      - "Conference/Seminar Name" 
      - "Month Year"
      - "City, Country"
      - "Conference Talk" # or "Seminar", "Poster", etc.
    description: "Talk abstract..." # or summary, etc.
    links:
      - text: "Slides"
        file: "/path/to/slides.pdf" # or 
        url: "https://url-to-slides.com"
      - text: "Video"
        url: "https://youtube.com/..."
---
```

### Lecture Notes Page (lecture_notes.md)

Structure for course materials:

```yaml
---
layout: lecture_notes
title: "Lecture Notes"
permalink: /lecture-notes/
description: "Course materials and educational resources"

courses:
  - title: "Course Name"
    institution: "University Name"
    year: "2024"
    description: "Course description"
    topics:
      - "Topic 1"
      - "Topic 2"
    materials:
      - text: "Notes PDF"
        url: "/path/to/notes.pdf"
      - text: "Problem Sets"
        url: "/path/to/problems.pdf"
---
```

## Creating Blog Posts

Blog posts go in the `_posts/` directory with the naming convention: `YYYY-MM-DD-title.md`

```yaml
---
layout: post
title: "Your Post Title"
date: 2025-08-30
author: "Your Name"
excerpt: "Brief description for previews"
reading_time: 5  # Optional: estimated reading time in minutes
---

Your post content here using Markdown syntax.

## Math Support
The template includes MathJax for mathematical expressions:

Inline math: $E = mc^2$

Display math:
$$\int_{-\infty}^{\infty} e^{-x^2} dx = \sqrt{\pi}$$
```

## Customization Tips

### Colors & Styling

Modify CSS variables in `assets/css/style.css`:

```css
:root {
    --color-primary: #B8724D;        /* Your brand color */
    --color-primary-dark: #9A5F42;   /* Darker variant */
    --content-max-width: 700px;      /* Content width */
    --sidebar-width: 320px;          /* Sidebar width */
}
```

### Adding New Sections

1. Create a new layout in `_layouts/`
2. Add corresponding styles to `style.css`
3. Create the markdown file with appropriate front matter

### File Organization

Your Jekyll site should be organized as follows:

```
your-academic-website/
├── _config.yml              # Site configuration
├── _layouts/                # HTML templates for different page types
│   ├── default.html
│   ├── cv.html
│   ├── research.html
│   ├── talks.html
│   ├── lecture_notes.html
│   └── post.html
├── _posts/                  # Blog posts and articles
│   └── 2025-09-03-how-to-setup.md
├── assets/                  # Static files
│   ├── css/
│   │   └── style.css
│   ├── images/
│   │   └── profile.jpg
│   └── files/               # PDFs, slides, etc.
│       ├── cv.pdf
│       ├── thesis.pdf
│       └── slides/
├── index.md                 # Homepage
├── cv.md                    # CV page
├── research.md              # Research page
├── talks.md                 # Talks page
├── lecture_notes.md         # Lecture notes page
├── robots.txt              # Search engine instructions
├── Gemfile                 # Ruby dependencies
└── README.md               # Repository documentation
```

**Key Directories:**

- `_layouts/`: HTML templates for different page types
- `_posts/`: Blog posts and articles (named YYYY-MM-DD-title.md)
- `assets/`: Static files (CSS, images, PDFs)
- Root `.md` files: Main pages accessible via navigation

## Deployment

### GitHub Pages

1. Push your repository to GitHub
2. Go to Settings > Pages
3. Select "Deploy from a branch"
4. Choose "main" branch and "/ (root)" folder
5. Your site will be available at `https://username.github.io/repository-name`

## Troubleshooting

- **Build errors**: Check `_config.yml` syntax and front matter formatting
- **Styling issues**: Verify CSS file paths and variable names
- **Math not rendering**: Ensure MathJax configuration is correct in `default.html`
- **Images not loading**: Check file paths and case sensitivity

For more help, refer to the [Jekyll documentation](https://jekyllrb.com/docs/).
