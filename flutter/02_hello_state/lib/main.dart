import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        title: 'Hello State',
        theme: ThemeData(
          // This is the theme of your application.
          //
          // Try running your application with "flutter run". You'll see the
          // application has a blue toolbar. Then, without quitting the app, try
          // changing the primarySwatch below to Colors.green and then invoke
          // "hot reload" (press "r" in the console where you ran "flutter run",
          // or simply save your changes to "hot reload" in a Flutter IDE).
          // Notice that the counter didn't reset back to zero; the application
          // is not restarted.
          primarySwatch: Colors.blue,
        ),
        home: Scaffold(
          appBar: AppBar(
            title: Text("My first StatefulWidget"),
          ),
          body: Center(child: HelloText()),
        ));
  }
}

class HelloText extends StatefulWidget {
  const HelloText({Key? key}) : super(key: key);

  @override
  _HelloTextState createState() => _HelloTextState();
}

class _HelloTextState extends State<HelloText> {
  final myController = TextEditingController();
  String text = '';
  @override
  Widget build(BuildContext context) {
    return Column(children: [
      TextFormField(
        onChanged: (newText) => updateText(newText),
        controller: myController,
        decoration: const InputDecoration(
            labelText: 'What is your name?', border: UnderlineInputBorder()),
      ),
      Divider(),
      Text(text)
    ]);
  }

  void updateText(String newText) {
    setState(() {
      if (newText.isEmpty) {
        text = '😎';
      } else {
        text = '👋 Hello ' + newText.toUpperCase() + "!";
      }
    });
  }
}
