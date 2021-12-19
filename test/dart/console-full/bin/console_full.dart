import 'dart:convert';
import 'dart:developer';

import 'package:http/http.dart' as http;

final parameters = {"type": "login"};

final url = Uri.http("localhost:8081", "/signin", parameters);
final headers = {
  "Content-type": "application/json",
  "Accept": "application/json"
};
final payload = jsonEncode({
  "Email": "email7@c.org",
  "Password": "pwd7",
  "Firstname": "",
  "Lastname": ""
});

class Signin {
  final String email;
  final String password;
  final String fname;
  final String lname;
  final String id;

  Signin({
    required this.email,
    required this.password,
    required this.fname,
    required this.lname,
    required this.id,
  });

  factory Signin.fromJson(Map<String, dynamic> json) {
    return Signin(
        email: json["Email"],
        password: json["Password"],
        fname: json["Firstname"],
        lname: json["Lastname"],
        id: json["Id"]);
  }
}

class Customise {
  final String scheme;
  final String id;

  Customise({required this.scheme, required this.id});

  factory Customise.fromJson(Map<String, dynamic> json) {
    return Customise(scheme: json["Colorscheme"], id: json["Id"]);
  }
}

class Account {
  final Signin signin;
  final Customise cust;

  Account({required this.signin, required this.cust});

  factory Account.fromJson(Map<String, dynamic> json) {
    return Account(
      signin: Signin.fromJson(json["Signin"]),
      cust: Customise.fromJson(json["Cust"]),
    );
  }
}

class AResponse {
  final int code;
  final Account account;

  AResponse({
    required this.code,
    required this.account,
  });

  factory AResponse.fromJson(Map<String, dynamic> json) {
    return AResponse(
      code: json['Code'],
      account: Account.fromJson(json["Account"]),
    );
  }
}

AResponse fetchAResponse(String body) {
  return AResponse.fromJson(jsonDecode(body));
}

void req() async {
  http.Response response = await http.post(
    url,
    headers: headers,
    body: payload,
  );
  AResponse obj = fetchAResponse(response.body);
  print(obj.account.signin.fname);
}

void main() {
  req();
}
