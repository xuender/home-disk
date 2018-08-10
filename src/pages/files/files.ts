import { Component } from '@angular/core';
import { NavController, NavParams } from 'ionic-angular';
import { File } from '../../domain/file'
import { FilesProvider } from '../../providers/files/files'
import { PreviewProvider } from '../../providers/preview/preview';
import { includes, pull } from 'lodash-es'
@Component({
  selector: 'page-files',
  templateUrl: 'files.html',
})
export class FilesPage {
  types: string[] = []
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
  toggle(type: string) {
    if (includes(this.types, type)) {
      pull(this.types, type)
    } else {
      this.types.push(type)
    }
  }
  isSelect(type: string) {
    return includes(this.types, type)
  }
}
