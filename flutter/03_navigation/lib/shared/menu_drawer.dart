import 'package:flutter/material.dart';
import 'package:navigation/screens/one.dart';
import 'package:navigation/screens/three.dart';

import '../screens/two.dart';

class MenuDrawer extends StatelessWidget {
  const MenuDrawer({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Drawer(child: ListView(children: _buildMenuItems(context)));
  }

  List<Widget> _buildMenuItems(BuildContext context) {
    List<Widget> menuItems = [];

    final List<String> menuTitles = ['Uno', 'Dos', 'Tres'];
    menuItems.add(_menuHeader());
    for (var element in menuTitles) {
      Widget screen = Container();
      menuItems.add(ListTile(
          title: Text(element, style: const TextStyle(fontSize: 19)),
          onTap: () {
            switch (element) {
              case 'Uno':
                screen = const ScreenOne();
                break;
              case 'Dos':
                screen = const ScreenTwo();
                break;
              case 'Tres':
                screen = const ScreenThree();
                break;
            }
            Navigator.of(context).pop();
            Navigator.of(context)
                .push((MaterialPageRoute(builder: (context) => screen)));
          }));
    }
    return menuItems;
  }

  Widget _menuHeader() {
    return const DrawerHeader(
        child: Text('Navigation App',
            style: TextStyle(color: Colors.indigo, fontSize: 32)));
  }
}
