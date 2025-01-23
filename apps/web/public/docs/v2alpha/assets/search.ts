import { connectRefinementList } from 'instantsearch.js/es/connectors';
const { algoliasearch, instantsearch } = window;
const { autocomplete } = window['@algolia/autocomplete-js'];
const { createLocalStorageRecentSearchesPlugin } =
  window['@algolia/autocomplete-plugin-recent-searches'];
const { createQuerySuggestionsPlugin } =
  window['@algolia/autocomplete-plugin-query-suggestions'];

enum LocalSettings {
  Query = 'query',
  Updated = 'updated',
}

function getSvgContent(url: string): string {
  const request = new XMLHttpRequest();
  request.open('GET', url, false);
  request.send();
  if (request.status === 200) {
    return request.responseText;
  } else {
    return '';
  }
}

interface RefinementListItem {
  value: string;
  label: string;
  count: number;
  isRefined: boolean;
}

interface RenderOptions {
  items: RefinementListItem[];
  canRefine: boolean;
  refine: (item_value: string) => void;
  sendEvent: (eventType: string, facetValue: string) => void;
  createURL: (item_value: string) => string;
  isFromSearch: boolean;
  searchForItems: (query: string) => void;
  hasExhaustiveItems: boolean;
  isShowingMore: boolean;
  canToggleShowMore: boolean;
  toggleShowMore: () => void;
  widgetParams: object;
}

interface CustomRefinementListOptions {
  container: string;
  title: string;
  attribute: string;
}

function customRefinementList(options: CustomRefinementListOptions) {
  const { container, title, attribute } = options;

  function renderRefinementList() {
    return (renderOptions: RenderOptions, isFirstRender: boolean) => {
      const { items, canRefine, refine } = renderOptions;
      const htmlContainer = document.querySelector(container);

      if (items.length === 0) {
        htmlContainer.innerHTML = '';
        return;
      }

      if (!canRefine) {
        htmlContainer.innerHTML = '';
        return;
      }

      htmlContainer.innerHTML = `
      <h2>${title}</h2>
      <div class="ais-RefinementList">
        <ul class="ais-RefinementList-list">
          ${items
            .map(
              (item) => `
          <li class="ais-RefinementList-item${
            item.isRefined ? ' ais-RefinementList-item--selected' : ''
          }">
            <div>
              <label class="ais-RefinementList-label" data-value="${
                item.value
              }">
                <input type="checkbox" class="ais-RefinementList-checkbox" ${
                  item.isRefined ? 'checked ' : ''
                }value="${item.label}">
                <span class="ais-RefinementList-labelText">${item.label}</span>
                <span class="ais-RefinementList-count">${item.count}</span>
              </label>
            </div>
          </li>
          `
            )
            .join('')}
        </ul>
      </div>
      `;

      [...htmlContainer.querySelectorAll('label')].forEach((element) => {
        element.addEventListener('click', (event) => {
          event.preventDefault();
          refine(element.dataset.value);
        });
      });
    };
  }

  return connectRefinementList(renderRefinementList())({
    attribute: attribute,
  });
}

(function () {
  const currentScriptTag = document.currentScript;

  const dataset = currentScriptTag && currentScriptTag.dataset;
  const { appid, apikey, indexname } = dataset || {};

  const searchClient = algoliasearch(appid, apikey);
  const search = instantsearch({
    indexName: indexname,
    searchClient,
    insights: true,
    routing: {
      router: instantsearch.routers.history(),
      stateMapping: {
        routeToState(routeState) {
          localStorage.removeItem(LocalSettings.Updated);
          localStorage.setItem(LocalSettings.Query, routeState.q);
          const returnValue = {
            [indexname]: {
              query: routeState.q,
              page: routeState.page,
            },
          };
          return returnValue;
        },
        stateToRoute(uiState) {
          const indexUiState = uiState[indexname];
          const returnValue = {
            page: indexUiState.page,
            q: indexUiState.query,
          };
          return returnValue;
        },
      },
    },
  });

  const virtualSearchBox = instantsearch.connectors.connectSearchBox(() => {
    const inputBox =
      (document.getElementById('autocomplete-0-input') as HTMLInputElement) ||
      null;
    if (inputBox) {
      const query = localStorage.getItem(LocalSettings.Query);
      if (localStorage.getItem(LocalSettings.Updated)) {
        return;
      }
      if (query !== 'undefined') {
        inputBox.value = query;
      }
      localStorage.setItem(LocalSettings.Updated, 'yes');
    }
  });

  search.addWidgets([
    virtualSearchBox({}),

    instantsearch.widgets.hits({
      container: '#hits',
      transformItems: function (items) {
        return items.map((item) => {
          const logo_url = `/icons/${item.logo_name}.svg`;
          const svgContent = getSvgContent(logo_url);

          return {
            ...item,
            logo: svgContent,
          };
        });
      },
      templates: {
        item: (hit, { html, components }) => {
          return `
            <article>
              <h1>
                <a class="ProductGrid--link" href="${hit.url}">
                  ${
                    typeof hit.type === 'string' && hit.type === 'system'
                      ? `<div>${hit.logo}</div>`
                      : ''
                  }
                  ${hit.title}
                </a>
              </h1>
              <ul class="breadcrumbs">
                ${hit.breadcrumbs
                  .map((breadcrumb) => {
                    return `
                    <li>
                      <a href="${breadcrumb.url}">${breadcrumb.title}</a>
                    </li>`;
                  })
                  .join('')}
              </ul>
              <p>
                ${hit.description === undefined ? '' : hit.description}
              </p>
            </article>
            `;
        },
      },
    }),

    customRefinementList({
      container: '#focusareas-list',
      title: 'Focus Areas',
      attribute: 'focusareas',
    }),

    customRefinementList({
      container: '#platform-contexts-list',
      title: 'Platform Contexts',
      attribute: 'context',
    }),

    customRefinementList({
      container: '#platform-systems-list',
      title: 'System Types',
      attribute: 'systems',
    }),

    instantsearch.widgets.configure({
      hitsPerPage: 8,
    }),

    instantsearch.widgets.pagination({
      container: '#pagination',
    }),

    instantsearch.widgets.stats({
      container: '#stats',
    }),
  ]);

  search.start();

  const recentSearchesPlugin = createLocalStorageRecentSearchesPlugin({
    key: 'instantsearch',
    limit: 3,
    transformSource({ source }) {
      return {
        ...source,
        onSelect({ setIsOpen, setQuery, item, event }) {
          onSelect({ setQuery, setIsOpen, event, query: item.label });
        },
      };
    },
  });

  const querySuggestionsPlugin = createQuerySuggestionsPlugin({
    searchClient,
    indexName: 'first_index',
    getSearchParams() {
      return recentSearchesPlugin.data.getAlgoliaSearchParams({
        hitsPerPage: 6,
      });
    },
    transformSource({ source }) {
      return {
        ...source,
        sourceId: 'querySuggestionsPlugin',
        onSelect({ setIsOpen, setQuery, event, item }) {
          onSelect({ setQuery, setIsOpen, event, query: item.query });
        },
        getItems(params) {
          if (!params.state.query) {
            return [];
          }

          return source.getItems(params);
        },
      };
    },
  });

  autocomplete({
    container: '#searchbox',
    openOnFocus: true,
    placeholder: 'Search docs...',
    detachedMediaQuery: 'none',
    onSubmit({ state }) {
      setInstantSearchUiState({ query: state.query });
    },
    plugins: [recentSearchesPlugin, querySuggestionsPlugin],
  });

  function setInstantSearchUiState(indexUiState) {
    search.mainIndex.setIndexUiState({ page: 1, ...indexUiState });
  }

  function onSelect({ setIsOpen, setQuery, event, query }) {
    if (isModifierEvent(event)) {
      return;
    }

    setQuery(query);
    setIsOpen(false);
    setInstantSearchUiState({ query });
  }

  function isModifierEvent(event) {
    const isMiddleClick = event.button === 1;

    return (
      isMiddleClick ||
      event.altKey ||
      event.ctrlKey ||
      event.metaKey ||
      event.shiftKey
    );
  }
})();
