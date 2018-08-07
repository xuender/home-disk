import { Component } from '@angular/core';

import { AboutPage } from '../about/about';
import { FilesPage } from '../files/files';
import { HomePage } from '../home/home';
import { FilesProvider } from '../../providers/files/files';

@Component({
  templateUrl: 'tabs.html'
})
export class TabsPage {

  tab1Root = HomePage;
  tab2Root = AboutPage;
  tab3Root = FilesPage

  constructor(private filesProvider: FilesProvider) {

  }
  reset(e: any) {
    if (e.index == 1) {
      this.filesProvider.reset()
    }
  }
}
