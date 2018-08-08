import { Component } from '@angular/core';
import { NavController } from 'ionic-angular';
import { FilesProvider } from '../../providers/files/files';
import { File } from '../../domain/file'

@Component({
  selector: 'page-upload',
  templateUrl: 'upload.html',
})
export class UploadPage {
  constructor(
    public navCtrl: NavController,
    private filesProvider: FilesProvider
  ) {
  }
  onUpload(f: File) {
    this.filesProvider.news.push(f)
  }
}
