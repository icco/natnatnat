{{ template "includes/header" "Search" }}

<form>
  <input type="text" name="s" class="input-text" />
  <input type="submit" />
</form>

{{.Count}} Results

{{ template "includes/footer" }}
