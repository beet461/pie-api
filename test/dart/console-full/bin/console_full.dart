import 'dart:convert';

import 'package:http/http.dart' as http;

void req() async {
  Uri url = Uri.http("localhost:8081", "/customisation");
  String payload = jsonEncode({
    'Colorscheme': "",
    'Id': "htP1whqOPcQb5CIdr6qV58A49sDD834teQ2B",
  });
  http.Response response = await http.post(url, body: payload);
  print(response.body);
}

void main() {
  req();
}
