import { Component } from '@angular/core';
import { NavController, NavParams } from 'ionic-angular';
import { File } from '../../domain/file'
import { FilesProvider } from '../../providers/files/files'
import { PreviewProvider } from '../../providers/preview/preview';

@Component({
  selector: 'page-files',
  templateUrl: 'files.html',
})
export class FilesPage {
  constructor(
    public navCtrl: NavController,
    public navParams: NavParams,
    public files: FilesProvider,
    private preview: PreviewProvider
  ) {
  }
  doInfinite(infiniteScroll) {
    this.files.loadDay()
      .then(r => infiniteScroll.complete())
  }
  onSelectFile(file: File) {
    this.preview.preview(file)
  }
}
