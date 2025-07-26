import * as events from './events';
import * as contents from './contents';
import * as timeago from './timeago';
import * as navs from './navlinks';

import { autocomplete } from '@algolia/autocomplete-js';
import { createLocalStorageRecentSearchesPlugin } from '@algolia/autocomplete-plugin-recent-searches';
import { createQuerySuggestionsPlugin } from '@algolia/autocomplete-plugin-query-suggestions';
import algoliasearch from 'algoliasearch/lite';
import instantsearch from 'instantsearch.js/dist/instantsearch.production.min';

declare global {
    interface Window {
        algoliasearch: typeof algoliasearch;
        instantsearch: typeof instantsearch;
        '@algolia/autocomplete-js': {
            autocomplete: typeof autocomplete;
        };
        '@algolia/autocomplete-plugin-recent-searches': {
            createLocalStorageRecentSearchesPlugin: typeof createLocalStorageRecentSearchesPlugin;
        };
        '@algolia/autocomplete-plugin-query-suggestions': {
            createQuerySuggestionsPlugin: typeof createQuerySuggestionsPlugin;
        };
    }
}

(function () {
    navs.init();
    timeago.init();
    events.load();
    events.focus();
    events.mobile();
    events.dropdowns();
    events.clipboardButton();
    events.copy();
    contents.toc();
    events.toggleSidebar();
    events.activeTab();
    events.tabs();
})();
