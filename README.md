# Jekyll Academic Website Template

A clean Jekyll template for academics and researchers. Features structured templates for CV, publications, talks, and course materials.

## Quick Start

### Prerequisites

- Ruby (2.7+)
- Bundler gem
- Git

### Installation

```bash
git clone https://github.com/paul-the-wizord/website.git
cd your-repo-name
bundle install
bundle exec jekyll serve
```

Change the configurations (see below).

Visit `http://localhost:4000` to see your site.

## Configuration

Edit `_config.yml` with your information:

```yaml
title: "Your Name"
email: "your.email@example.com"
baseurl: "/your-repo-name"
url: "https://yourusername.github.io"
github_username: "your-github-username"
baseurl: "" # For Github deployment set to repo-name for local deployment set to ""
```

Update content files (cv.md, research.md, talks.md, etc.) and replace the profile image in `/assets/images/profile.jpg` as well as the favicon files.

## Deployment

Push to GitHub and enable Pages in repository settings. Your site will be available at `https://yourusername.github.io/repository-name`.

## Documentation

For complete setup instructions, customization options, and usage examples, visit: [https://paul-the-wizord.github.io/website/post/2025-09-03-website-setup-guide/](https://paul-the-wizord.github.io/website/post/2025-09-03-website-setup-guide/)

## License

This project uses a dual licensing approach:

- **Template & Code**: Licensed under the MIT License - see [LICENSE-CODE](LICENSE-CODE)
- **Personal Content**: All rights reserved - see [LICENSE-CONTENT](LICENSE-CONTENT)

The Jekyll template, layouts, and styling may be freely used. Personal content
(CV, research info, photos) requires permission.
