<a name="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<br />
<div align="center">
  <h3 align="center">Diade</h3>
  <p align="center">
    Compilation of international celebrations
    <br />
    <a href="https://github.com/rfdez/diade"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/rfdez/diade">View Demo</a>
    ·
    <a href="https://github.com/rfdez/diade/issues">Report Bug</a>
    ·
    <a href="https://github.com/rfdez/diade/issues">Request Feature</a>
  </p>
</div>

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

## About The Project

Compilation of international celebrations. At this moment the project only supports spanish language.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

[![Go][Go.dev]][Go-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* docker
* docker compose
* go

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/rfdez/diade.git
   ```
2. Start the project
   ```sh
   docker-compose up -d
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Usage

The API server will be running on port 8080. You can access the API celebrations endpoint at http://localhost:8080/celebrations. You can also access the API celebrations endpoint at http://localhost:8080/celebrations?date=2021-01-01.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

[contributors-shield]: https://img.shields.io/github/contributors/rfdez/diade.svg?style=for-the-badge
[contributors-url]: https://github.com/rfdez/diade/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/rfdez/diade.svg?style=for-the-badge
[forks-url]: https://github.com/rfdez/diade/network/members
[stars-shield]: https://img.shields.io/github/stars/rfdez/diade.svg?style=for-the-badge
[stars-url]: https://github.com/rfdez/diade/stargazers
[issues-shield]: https://img.shields.io/github/issues/rfdez/diade.svg?style=for-the-badge
[issues-url]: https://github.com/rfdez/diade/issues
[license-shield]: https://img.shields.io/github/license/rfdez/diade.svg?style=for-the-badge
[license-url]: https://github.com/rfdez/diade/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/raul-fernandez-fernandez
[product-screenshot]: images/screenshot.png
[Go.dev]: https://img.shields.io/badge/go-000000?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
