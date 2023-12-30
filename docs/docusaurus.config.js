// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/vsLight');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Toae YaraHunter',
  tagline: 'Malware scanner for cloud-native, as part of CI/CD and at Runtime',
  url: 'https://yarahunter.local',
  baseUrl: '/',

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'throw',
  favicon: '/img/toae.png',

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/Sam12121/YaraHunter/docs/',
          breadcrumbs: true,
          routeBasePath: '/',
          // options for remark-admonitions
          admonitions: {}, 
        },
        blog: false,
        theme: {
          customCss: require.resolve('./static/css/toae.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      navbar: {
        hideOnScroll: false,
        title: '',
        logo: {
          alt: 'Toae',
          src: '/img/toae-logo-black.svg',
          srcDark: '/img/toae-logo-white.svg',
        },
        items: [
          {
            type: 'doc',
            docId: 'yarahunter/index',
            label: 'YaraHunter Documentation',
          },
          { 
            to: 'https://toae.io/',
            label: 'toae.io',
            position: 'right',
          },
          {
            href: 'https://github.com/Sam12121',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        logo: {
          alt: 'Toae, Inc',
          src: '/img/toae-logo-white.svg',
          width: 160,
          height: 51,
        },
        copyright: `Copyright © ${new Date().getFullYear()} Toae, Inc. Built with Docusaurus.`,
      },
      colorMode: {
        defaultMode: 'light',
        disableSwitch: false,
        respectPrefersColorScheme: true,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
};

module.exports = config;
