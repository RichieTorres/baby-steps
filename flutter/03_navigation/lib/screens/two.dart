import 'package:flutter/material.dart';
import 'package:navigation/shared/menu_drawer.dart';

import '../shared/menu_bottom.dart';

class ScreenTwo extends StatelessWidget {
  const ScreenTwo({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Screen Two'),
      ),
      drawer: const MenuDrawer(),
      bottomNavigationBar: const MenuBottom(),
      body: Column(children: const [
        Center(
            child: Text(
          '2',
          style: TextStyle(fontSize: 20),
        )),
        Center(child: Text('This is the second view'))
      ]),
    );
  }
}
