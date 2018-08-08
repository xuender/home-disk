import { Component } from '@angular/core';
import { Platform } from 'ionic-angular';

import { AboutPage } from '../about/about';
import { FilesPage } from '../files/files';
import { FilesProvider } from '../../providers/files/files';
import { UploadPage } from '../upload/upload';

@Component({
  templateUrl: 'tabs.html'
})
export class TabsPage {

  tab1Root = UploadPage
  tab2Root = AboutPage;
  tab3Root = FilesPage
  selectedIndex = 0

  constructor(public filesProvider: FilesProvider,
    private plt: Platform
  ) {
    if(!this.plt.is('mobile') && !this.plt.is('mobileweb')){
      this.selectedIndex = 2
    }
  }
}
