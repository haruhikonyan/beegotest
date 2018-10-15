<form action="/post" method="post">
  <div class="form-group">
    <label>タイトル</label>
    <input name="title" class="form-control" placeholder="タイトル">
  </div>
  <div class="form-group">
    <label>内容</label>
    <textarea name="body" class="form-control" rows="3" placeholder="書き込む内容"></textarea>
  </div>
  <button type="submit" class="btn btn-primary">送信</button>
</form>

<hr>

{{range $val := .posts}}
  <div class="card bg-light mb-3">
    <div class="card-header">{{$val.Id}}. ななしさん</div>
    <div class="card-body">
      <h5 class="card-title">{{$val.Title}} </h5>
      <p class="card-text">{{$val.Body}}</p>
    </div>
  </div>
{{end}}