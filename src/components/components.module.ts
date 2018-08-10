import { NgModule } from '@angular/core';
import { IonicModule } from 'ionic-angular'
import { FileUploadModule } from 'ng2-file-upload'
//import { NgxUploaderModule } from 'ngx-uploader'
import { IonicImageViewerModule } from 'ionic-img-viewer'

import { HdUploadComponent } from './hd-upload/hd-upload';
import { HdFilesComponent } from './hd-files/hd-files';
@NgModule({
  declarations: [
    HdUploadComponent,
    HdFilesComponent,
  ],
  imports: [
    IonicModule,
    FileUploadModule,
    IonicImageViewerModule,
  ],
  exports: [
    HdUploadComponent,
    FileUploadModule,
    HdFilesComponent,
  ]
})
export class ComponentsModule { }
