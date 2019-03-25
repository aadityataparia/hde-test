# HDE Design test

## Overview 

__Mission 1__: Design the UI using the given images
Assignment: Create the UI mockup using HTML, CSS, Javascript, etc. based on the images created by our designer (see images.zip).

__Mission 2__: Extend the UI in order to implement a new feature
Assignment: In the search result section (see result.png), extend the UI mockup so that user can inspect the body of each emails from the search result.

__Rules__:
- Please use the images that we provide for the assignments.
- Please produce high-fidelity UI mockup for the assignments.
- Any fixed sample texts can be used as the data within the UI.
- For Mission 2: The extended design should allow user to inspect multiple email bodies at once.

## Basic requirements

- Show email results (From, To, Subject and date)
- Filter email results based on dates
- Order email results by date
- Show mail archiver logo when no results are there
- __Challenge__ Add email body to results shown
  - Show bodies of all results in same page (trucated)
  - Ability to view complete email body on same page
- Use mock up images given to achieve this

## Frameworks/libraries used

- [sifrr-dom](https://github.com/sifrr/sifrr/tree/master/packages/browser/sifrr-dom) Sifrr-dom is a frontend framework like react, which is built on top of [CustomElements](https://developer.mozilla.org/en-US/docs/Web/Web_Components/Using_custom_elements) and handles states, one/two way binding, event listeners etc. I am author of this framework.
- NodeJS/Express for hosting static server

## App structure (Mission 1)

![result](https://user-images.githubusercontent.com/13755349/54894903-3fc54d00-4f00-11e9-866b-73e88f8dc3f6.png)

### Main Container 

Main container is responsible for getting the email results to be shown and filter them based on dates given in search bar.

Custom element tag: `main-element`

### Search Bar

Gives `from` and `to` date to filter the results

Included in `main-element`, has two way data binding with `from` and `to` dates, to filter the results. And one event listener for search button which triggers the filtering.

### Results Table

Shows result row for each result, and orders results by date, shows logo when no results are present

container tag name: `result-container` shows result-table when results are present, else shows logo
results table: `result-table` 

### Result Row

Shows one result

one result tag name: `result-row`, handles data bindings for `from`, `to`, `subject`, `date`

## Data handling

Since there was no backend to provide data, I wrote a simple random data generator, which would give random email results with from, to, subject, date. Random data is generate in `main-element` and filtered based on `from` `to` dates given in search bar.

This filtered data is passed to `result-container` which shows logo if results array length is 0 and shows `result-table` when results array length is not 0. It shown number of results in top as well. It then passes the results that need to be shown to `result-table`. 

`result-table` then show headings and one `result-row` for each result, arranged by date DESC. When you click `date` in heading, it toggles the arrangement between ASC & DESC and changes the arrow sign shown. It updates results with sifrr's keyed implementation, based on unique id for each row.

`result-row` has four columns for showing `from` `to` `subject` `date`.

## Implementing Mission 2 (adding body to result)

Since we already have data flow established, we just need to add body to the result data. I wrote a random body generator which generates random sentences for body. We just pass the `body` as well with the results data we were passing already. Only thing left is showing the `body` in `result-row` table.

To show the table I added a new row in `result-row` element, which shows three lines of the body. To show full body, I added a click event listener on body row, which toggles it between showing just 3 rows and showing full body. 

## Handling CSS

Since we are using custom elements and dom is added under shadow-root (except `result-row`) all the CSSs defined in an element is scoped to that element only. Hence, I defined needed CSS in individual elements, rather than having a different CSS file. 

## How SIfrr-dom Works

- You define needed custom elements which have `template` and class which extends `Sifrr.Dom.Element` and register it using `Sifrr.Dom.register(className)`. You define data bindings using `${}` directives in html and everything else is handled by sifrr-dom. You can reference current element in binding with `this`.

- `ResultRow` class will be registered as `result-row`

- Folder structure: for an sifrr element with tag `result-row`, you will have a file `element/result/row.(html|js)`

- Elements are loaded asynchronously using `Sifrr.Dom.load('result-row')`, and all the elements load elements that they need.

- Event listeners can be added on elements (eg. `_click=${function}`)
